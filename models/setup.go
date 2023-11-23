package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDb() {
	db, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/letstalk?parseTime=true"))
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Messege{})

	DB = db
}
