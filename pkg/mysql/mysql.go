package mysql

import (
	"fmt"
	"git.amocrm.ru/dmiroshnikov/unisender_integration/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	DB *gorm.DB
}

func NewDatabase(sql config.MySQL) Database {
	USER := sql.DBUser
	PASS := sql.DBPassword
	HOST := sql.DBHost
	DBNAME := sql.DBName

	URL := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", USER, PASS,
		HOST, DBNAME)
	db, err := gorm.Open(mysql.Open(URL))

	if err != nil {
		panic("Failed to connect to database!")

	}
	fmt.Println("Database connection established")
	return Database{
		DB: db,
	}
}
