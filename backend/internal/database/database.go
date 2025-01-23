package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"gmxBackend/config"
)

var (
	db *gorm.DB
)

func ConnectDB() error {
	err := CreateDBIfNotExists()
	if err != nil {
		return err
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=3s",
		config.AppConfig.MySQL.DBUser,
		config.AppConfig.MySQL.DBPassword,
		config.AppConfig.MySQL.DBHost,
		config.AppConfig.MySQL.DBPort,
		config.AppConfig.MySQL.DBName)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: newLogger})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		return err
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get database instance: %v", err)
		return err
	}

	sqlDB.SetMaxIdleConns(25)
	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	log.Println("MySQL database connected successfully")
	return nil
}

func GetDB() *gorm.DB {
	return db
}

func CloseDB() error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	if err := sqlDB.Close(); err != nil {
		return err
	}

	return nil
}

// CreateDBIfNotExists creates the database if it doesn't exist
func CreateDBIfNotExists() error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/",
		config.AppConfig.MySQL.DBUser,
		config.AppConfig.MySQL.DBPassword,
		config.AppConfig.MySQL.DBHost,
		config.AppConfig.MySQL.DBPort)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("error opening database connection: %v", err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + config.AppConfig.MySQL.DBName)
	if err != nil {
		return fmt.Errorf("error creating database: %v", err)
	}

	log.Printf("Database '%s' created or already exists", config.AppConfig.MySQL.DBName)
	return nil
}

// DeleteDB deletes the database
func DeleteDB() error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/",
		config.AppConfig.MySQL.DBUser,
		config.AppConfig.MySQL.DBPassword,
		config.AppConfig.MySQL.DBHost,
		config.AppConfig.MySQL.DBPort)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("error opening database connection: %v", err)
	}
	defer db.Close()

	_, err = db.Exec("DROP DATABASE IF EXISTS " + config.AppConfig.MySQL.DBName)
	if err != nil {
		return fmt.Errorf("error deleting database: %v", err)
	}

	log.Printf("Database '%s' deleted", config.AppConfig.MySQL.DBName)
	return nil
}
