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


	r.GET("/signin", controller.SigninGet)
	r.POST("/signin",controller.SiginPost)
	r.GET("/index",controller.HomeGet)

	r.GET("/loginsusscess",middleware.AuthMiddleware(),controller.LoginSuccess)
	r.GET("/login",controller.LoginGet)
	r.POST("/login",controller.LoginPost)



	r.Run(":8080")
}
