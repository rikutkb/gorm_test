package main

import (
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("cannot load .env")
	}

	Host = "mysql_host"
	Port = "13306"
	DBName = os.Getenv("MYSQL_DATABASE")
	Password = os.Getenv("MYSQL_PASSWORD")
	UserName = os.Getenv("MYSQL_USER")
	dsn := UserName + ":" + Password + "@" + "tcp(" + Host + ":" + Port + ")/" + DBName + "?charset=utf8mb4&parseTime=True&loc=Local"
	gormDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}
}
