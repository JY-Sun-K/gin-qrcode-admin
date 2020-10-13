package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"qcode/controller"
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
	//r.Use(middleware.Cors())
	//r.Use(middleware.JwtVerify)
	userv1:=r.Group("/user")
	{
		userv1.GET("/signin", controller.SigninGet)
		userv1.POST("/signin",controller.SiginPost)


		userv1.GET("/loginsusscess",middleware.AuthMiddleware(),controller.LoginSuccess)
		userv1.GET("/login",controller.LoginGet)
		userv1.POST("/login",controller.LoginPost)
	}
	admin:=r.Group("/rootuser")
	{
		admin.GET("/index",controller.HomeGet)
		admin.POST("/users/deluser",controller.DelUserPost)
		admin.GET("/users/deluser",controller.DelUserGet)
		admin.GET("/user",controller.VisitUser)
	}






	r.Run(":8080")
}
