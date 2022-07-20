package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"os"
)

type Database struct {
	Host string
	Port string
	DatabaseName string
	Username string
	Password string
}


func LoadEnv() Database {
	dbHost := os.Getenv("DB_HOST")
	dbPort:= os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_DATABASE")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	cfg := Database{
		Host : dbHost,
		Port  : dbPort,
		Username :dbUsername,
		Password : dbPassword,
		DatabaseName : dbName,
	}
	return cfg
}

func InitDB(cfg Database) *gorm.DB  {
	dbString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",cfg.Username,cfg.Password, cfg.Host,cfg.Port,cfg.DatabaseName)
	db, _ := gorm.Open(postgres.Open(dbString), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	return db
}


func GetDBConnection() *gorm.DB {
	loadEnv := LoadEnv()
	dbConnect := InitDB(loadEnv)
	return dbConnect
}


