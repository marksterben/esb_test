package server

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	username = viper.GetString("DB_USERNAME")
	password = viper.GetString("DB_PASSWORD")
	host     = viper.GetString("DB_HOST")
	port     = viper.GetString("DB_PORT")
	database = viper.GetString("DB_DATABASE")
)

func DB() *gorm.DB {
	dsn := fmt.Sprintf("%s%s@tcp(%s%s)/%s", username, password, host, port, database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Unable to connect to database. Err:", err)
	}

	return db
}
