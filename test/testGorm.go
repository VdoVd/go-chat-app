package main

import (
	"chat/models"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/ginchat?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.UserBasic{})

	user := &models.UserBasic{}
	user.Name = "申专"
	db.Create(user)

	fmt.Println(db.First(user, 1))

	db.Model(user).Update("PassWord", "1234-")
}
