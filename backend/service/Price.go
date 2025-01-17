package service

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"gmxBackend/models"
	"gmxBackend/repository"
	"sync/atomic"
)

type PriceService struct {
	priceBuffer *models.PriceBuffer
	priceRepo   *repository.PriceRepository
}

func NewPriceService(_priceBuffer *models.PriceBuffer, _priceRepo *repository.PriceRepository) *PriceService {
	return &PriceService{
		priceBuffer: _priceBuffer,
		priceRepo:   _priceRepo,
	}
}

func (p *PriceService) UpdatePrices(updates []models.Price) error {
	activeIdx := p.priceBuffer.ActiveIndex.Load().(int)
	// 计算非活跃缓冲区索引
	inactiveIdx := 1 - activeIdx

	// 获取非活跃缓冲区
	inactiveBuffer := p.priceBuffer.Buffers[inactiveIdx]

	// 复制当前活跃缓冲区的数据到非活跃缓冲区
	for k, v := range p.priceBuffer.Buffers[activeIdx].Prices {
		inactiveBuffer.Prices[k] = v
	}

	// 更新非活跃缓冲区的价格
	for _, update := range updates {
		inactiveBuffer.Prices[update.Symbol] = update
	}

	// 原子性地切换活跃缓冲区
	p.priceBuffer.ActiveIndex.Store(inactiveIdx)

	// 更新计数器
	atomic.AddUint64(&p.priceBuffer.UpdateCount, 1)

	p.priceRepo.UpdatePrice(updates)
	return nil
}

const (
	// OKX API配置
	apiKey     = ""
	secretKey  = ""
	passphrase = "Shi~1234"
	baseURL    = "https://www.okx.com"
)

// 价格响应结构体
type TickerResponse struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data []struct {
		InstID    string `json:"instId"`
		Last      string `json:"last"`
		BidPx     string `json:"bidPx"`
		AskPx     string `json:"askPx"`
		Volume24h string `json:"vol24h"`
	} `json:"data"`
}

// 生成签名
func generateSignature(timestamp, method, requestPath string) string {
	message := timestamp + method + requestPath
	mac := hmac.New(sha256.New, []byte(secretKey))
	mac.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(mac.Sum(nil))
}

func main() {
	// 设置请求路径
	endpoint := "/api/v5/market/ticker?instId=BTC-USDT"

	// 获取时间戳
	timestamp := fmt.Sprintf("%d", time.Now().UTC().Unix())

	// 创建请求
	client := &http.Client{}
	req, err := http.NewRequest("GET", baseURL+endpoint, nil)
	if err != nil {
		fmt.Printf("创建请求失败: %v\n", err)
		return
	}

	// 添加请求头
	sign := generateSignature(timestamp, "GET", endpoint)
	req.Header.Add("OK-ACCESS-KEY", apiKey)
	req.Header.Add("OK-ACCESS-SIGN", sign)
	req.Header.Add("OK-ACCESS-TIMESTAMP", timestamp)
	req.Header.Add("OK-ACCESS-PASSPHRASE", passphrase)

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("请求失败: %v\n", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("读取响应失败: %v\n", err)
		return
	}

	// 解析JSON响应
	var tickerResp TickerResponse
	if err := json.Unmarshal(body, &tickerResp); err != nil {
		fmt.Printf("解析JSON失败: %v\n", err)
		return
	}

	// 打印结果
	if len(tickerResp.Data) > 0 {
		fmt.Printf("BTC/USDT 最新价格: %s\n", tickerResp.Data[0].Last)
		fmt.Printf("买一价: %s\n", tickerResp.Data[0].BidPx)
		fmt.Printf("卖一价: %s\n", tickerResp.Data[0].AskPx)
		fmt.Printf("24小时成交量: %s\n", tickerResp.Data[0].Volume24h)
	} else {
		fmt.Println("未获取到价格数据")
	}
}
