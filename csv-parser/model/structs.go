package model

import (
	"time"

	"gorm.io/gorm"
)

type AmexTransaction struct {
	gorm.Model
	Date   time.Time
	Payee  string
	Amount string
}

type DBConfig struct {
	Host     string
	Username string
	Password string
	DBName   string
	Port     string
	SslMode  string
	Schema   string
}
