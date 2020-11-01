package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"qcode/controller/root"
	"qcode/controller/user"
	"qcode/middleware"
	"qcode/models"
)

func main() {

	db,err :=models.InitDB()
	if err!=nil {
		log.Fatalf("db connect err :%v",err.Error())
	}
	defer db.Close()

	r:=gin.Default()
	r.LoadHTMLGlob("views/*")
	r.StaticFS("/static",http.Dir("./static"))
	//r.Static("/static", "./static")
	//r.Use(middleware.Cors())
	//r.Use(middleware.JwtVerify)
	userv1:=r.Group("/user")
	{
		userv1.GET("/signin", user.SigninGet)
		userv1.POST("/signin", user.SiginPost)


		userv1.GET("/loginsusscess",middleware.AuthMiddleware(), user.LoginSuccess)
		userv1.GET("/login", user.LoginGet)
		userv1.POST("/login", user.LoginPost)
	}
	admin:=r.Group("/rootuser")
	{
		admin.GET("/index",middleware.RootAuthMiddleware(), root.HomeGet)
		admin.POST("/users/deluser",middleware.RootAuthMiddleware(), root.DelUserPost)
		admin.GET("/users/deluser",middleware.RootAuthMiddleware(), root.DelUserGet)
		admin.GET("/user/:id",middleware.RootAuthMiddleware(), root.VisitUser)
		admin.GET("/signin", root.SigninGet)
		admin.POST("/signin", root.SiginPost)



		admin.GET("/login", root.LoginGet)
		admin.POST("/login", root.LoginPost)
	}






	r.Run(":8080")
}
