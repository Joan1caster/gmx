package rabbitmq

import (
	"fmt"
	"sync"

	"github.com/streadway/amqp"
)

// 全局连接
var (
	conn      *amqp.Connection
	connOnce  sync.Once
	connMutex sync.Mutex
)

// 生产者和消费者、交换机管理
var (
	Producers = make(map[string]*Producer)
	Consumers = make(map[string]*Consumer)
	prodMutex sync.RWMutex
	consMutex sync.RWMutex
)

// 初始化连接
func InitConnection(url string) error {
	var err error
	connOnce.Do(func() {
		connMutex.Lock()
		defer connMutex.Unlock()
		conn, err = amqp.Dial(url)
	})

	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()
	err = ch.ExchangeDeclare(
		"default",
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	return err
}

// 获取连接
func GetConnection() *amqp.Connection {
	return conn
}

// Producer 结构体
type Producer struct {
	channel *amqp.Channel
	name    string
}

// Consumer 结构体
type Consumer struct {
	channel *amqp.Channel
	name    string
	queue   string
}

// 创建生产者
func CreateProducer(name string) (*Producer, error) {
	prodMutex.Lock()
	defer prodMutex.Unlock()

	if p, exists := Producers[name]; exists {
		return p, nil
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	producer := &Producer{
		channel: ch,
		name:    name,
	}
	Producers[name] = producer
	return producer, nil
}

// 获取生产者
func GetProducer(name string) (*Producer, error) {
	prodMutex.RLock()
	defer prodMutex.RUnlock()

	if p, exists := Producers[name]; exists {
		return p, nil
	}
	return nil, fmt.Errorf("producer %s not found", name)
}

// 发布消息
func (p *Producer) Publish(message []byte, ex_name string, ex_type string, routing_key string) error {
	// 发布消息
	return p.channel.Publish(
		ex_name,     // 交换机
		routing_key, // 路由键
		false,       // mandatory
		false,       // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        message,
		},
	)
}

// 创建消费者
func CreateConsumer(name string) (*Consumer, error) {
	consMutex.Lock()
	defer consMutex.Unlock()

	if c, exists := Consumers[name]; exists {
		return c, nil
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	consumer := &Consumer{
		channel: ch,
		name:    name,
	}
	Consumers[name] = consumer
	return consumer, nil
}

// 获取消费者
func GetConsumer(name string) (*Consumer, error) {
	consMutex.RLock()
	defer consMutex.RUnlock()

	if c, exists := Consumers[name]; exists {
		return c, nil
	}
	return nil, fmt.Errorf("consumer %s not found", name)
}

// 绑定队列到交换机
func (c *Consumer) BindQueue(ex_name string, routing_key string, queue_name string) error {
	// 声明队列
	_, err := c.channel.QueueDeclare(
		queue_name, // 队列名称
		true,       // 持久化
		false,      // 自动删除
		false,      // 排他性
		false,      // 等待确认
		nil,        // 参数
	)
	if err != nil {
		return err
	}

	// 绑定队列到交换机
	return c.channel.QueueBind(
		queue_name,  // 队列名称
		routing_key, // 路由键
		ex_name,     // 交换机
		false,       // 等待确认
		nil,         // 参数
	)
}

// 开始消费
func (c *Consumer) Consume(handler func([]byte) error, queue_name string) error {
	msgs, err := c.channel.Consume(
		queue_name, // 队列
		"",         // 消费者
		true,       // 自动确认
		false,      // 排他性
		false,      // 不等待
		false,      // 参数
		nil,
	)
	if err != nil {
		return err
	}

	go func() {
		for msg := range msgs {
			if err := handler(msg.Body); err != nil {
				fmt.Printf("Error handling message: %v\n", err)
			}
		}
	}()

	return nil
}

// 关闭生产者
func (p *Producer) Close() error {
	prodMutex.Lock()
	defer prodMutex.Unlock()
	delete(Producers, p.name)
	return p.channel.Close()
}

// 关闭消费者
func (c *Consumer) Close() error {
	consMutex.Lock()
	defer consMutex.Unlock()
	delete(Consumers, c.name)
	return c.channel.Close()
}

// 关闭所有生产者和消费者
func Shutdown() error {
	// 关闭所有生产者
	for _, p := range Producers {
		if err := p.Close(); err != nil {
			return err
		}
	}

	// 关闭所有消费者
	for _, c := range Consumers {
		if err := c.Close(); err != nil {
			return err
		}
	}

	// 关闭连接
	if conn != nil {
		return conn.Close()
	}
	return nil
}
