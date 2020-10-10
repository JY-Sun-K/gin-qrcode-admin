package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

)

type User struct {
	gorm.Model
	Name string `json:"name"`
	Password string `json:"password"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Adderss string `json:"adderss"`
}

var DB *gorm.DB

func InitDB()(*gorm.DB ,error ) {
	db,err := gorm.Open("mysql", "root:sjy1999@/qcode?charset=utf8&parseTime=True&loc=Local")
	DB=db
	db.AutoMigrate(&User{})
	return db,err
}

func GetDB() *gorm.DB {
	return DB
}

func GetUsers() []User {
	var users []User
	DB.Find(&users)
	return users
}