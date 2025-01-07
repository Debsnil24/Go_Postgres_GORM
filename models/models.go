package models

import "gorm.io/gorm"

type Repository struct {
	DB *gorm.DB
}

type Book struct {
	Author    string `json:"author"`
	Title     string `json:"title"`
	Publisher string `json:"publisher"`
}

type Books struct {
	ID			uint		`gorm:"primary key;autoIncrement" json:"id"`
	Author		*string		`json:"author"`
	Title		*string		`json:"title"`
	Publisher	*string		`json:"publisher"`
}