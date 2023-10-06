package tables

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	
)

var dsn = "root:apparattum@tcp(host.docker.internal)/Logbook?charset=utf8mb4&parseTime=True&loc=Local"
//127.0.0.1:3306 changed for --> host.docker.internal, so we can run it directly in the container
var db *gorm.DB

func ConectDB() {
	connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	fmt.Println("You're in!")
	db = connection
}

func Close() {
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	sqlDB.Close()
}

func GetDB() *gorm.DB {
	return db
}