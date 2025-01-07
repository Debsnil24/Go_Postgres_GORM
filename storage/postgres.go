package storage

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host		string
	Port		string
	Password	string
	User		string
	DBName		string
	SSLMode		string
}

func NewConnection(config *Config)(*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host,config.Port,config.User,config.Password,config.DBName,config.SSLMode)
	log.Println(dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Failed to connect to Database")
		return nil, err
	}
	log.Println("Database Connection Established susccessfully")
	return db, nil
}