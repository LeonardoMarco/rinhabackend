package config

import (
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
		Host:     "localhost",
		Port:     "5432",
		User:     "postgres",
		Password: "postgres",
		Dbname:   "rinhabackend",
	}

	db, err := gorm.Open("postgres",
		"host="+connection.Host+" port="+connection.Port+
			" user="+connection.User+" dbname="+connection.Dbname+
			" password="+connection.Password+" sslmode=disable")
	if err != nil {
		panic("Failed to connect database")
	}

	db.LogMode(false)

	return db
}
