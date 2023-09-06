package config

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Connection struct {
	Host     string
	Port     string
	User     string
	Password string
	Dbname   string
}

func ConnectDatabase() *gorm.DB {
	connection := Connection{
		Host:     os.Getenv("RINHA_DB_HOST"),
		Port:     os.Getenv("RINHA_DB_PORT"),
		User:     os.Getenv("RINHA_DB_USER"),
		Password: os.Getenv("RINHA_DB_PASSWORD"),
		Dbname:   os.Getenv("RINHA_DB_NAME"),
	}

	db, err := gorm.Open("postgres",
		"host="+connection.Host+" port="+connection.Port+
			" user="+connection.User+" dbname="+connection.Dbname+
			" password="+connection.Password+" sslmode=disable")
	if err != nil {
		panic("Failed to connect database")
	} else {
		fmt.Println("Conectou")
	}

	db.LogMode(true)

	return db
}
