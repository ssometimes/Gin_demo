package common

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDb() *gorm.DB {
	//driverName := "mysql"
	host := "localhost"
	port := "3306"
	database := "GinDemo"
	username := "root"
	password := "root"
	charset := "utf8"
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset)
	db, err := gorm.Open(mysql.New(mysql.Config{DSN: args}), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

func GetDb() *gorm.DB {
	DB = InitDb()
	return DB
}
