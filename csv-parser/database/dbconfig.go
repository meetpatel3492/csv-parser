package database

import (
	"csv-parser/model"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var Instance *gorm.DB
var err error

func Connect() {
	dbConfig := getDBConfig()
	connectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s ",
		dbConfig.Host,
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.DBName,
		dbConfig.Port,
		dbConfig.SslMode)

	Instance, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: fmt.Sprintf("%s.", dbConfig.Schema),
		},
	})
	if err != nil {
		log.Fatalf("Error creating database connection - %v", err)
	}
	log.Println("Connected to Database")
}

func Migrate(){
	Instance.AutoMigrate(&model.AmexTransaction{})
	log.Println("Database migration completed")
}

func getDBConfig() model.DBConfig {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading env variables %v", err)
	}
	return model.DBConfig{
		Host:     os.Getenv("HOSTNAME"),
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("PASSWORD"),
		DBName:   os.Getenv("DATABASE"),
		Port:     os.Getenv("PORT"),
		SslMode:  os.Getenv("SSL_MODE"),
		Schema:   os.Getenv("SCHEMA")}
}
