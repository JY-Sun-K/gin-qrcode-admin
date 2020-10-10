package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"qcode/models"
)

func HomeGet(ctx *gin.Context)  {
	users:=models.GetUsers()
	fmt.Println(users)
	ctx.HTML(http.StatusOK,"index.tmpl",gin.H{
		"code":200,
		"user":users,

	})
}

