package storage

import (
	"fmt"
	"log"

	"github.com/Debsnil24/Go_Postgres_GORM/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	models.Config
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