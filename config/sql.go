package config

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type DbConnection struct {
	Host     string
	Port     string
	DbName   string
	Username string
	Password string
	SslMode  string
}

func DbConfig() DbConnection {
	dbConfig := DbConnection{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		DbName:   os.Getenv("DB_NAME"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		SslMode:  os.Getenv("DB_SSL_MODE"),
	}

	return dbConfig
}

func GetDbInstance() *gorm.DB {
	dbConfig := DbConfig()

	dsn := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DbName,
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.SslMode,
	)
	fmt.Println(dsn)
	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		log.Fatalf(err.Error())
	}

	return db
}
