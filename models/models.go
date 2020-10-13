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

func CreateUser(name,hasedPassword,email,phone,adderss string){

	newuser := &User{
		Name:     name,
		Password: hasedPassword,
		Email:    email,
		Phone:    phone,
		Adderss:  adderss,
	}
	DB.Create(&newuser)
}

func DelUser(id string)  {
	DB.Where("id = ?", id).Delete(User{})
}

func SearchUserById(id string)User  {
	var user User
	DB.Where("id = ?", id).First(&user)
	return user
}