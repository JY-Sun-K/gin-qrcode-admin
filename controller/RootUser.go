package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"qcode/models"
	"qcode/pkg/qrcode"
)

func DelUserPost(ctx *gin.Context)  {

	userId:=ctx.PostForm("id")
	models.DelUser(userId)


}

func DelUserGet(ctx *gin.Context)  {
	users:=models.GetUsers()
	ctx.JSON(http.StatusOK,gin.H{
		"code":200,
		"users":users,
	})
}

func VisitUser(ctx *gin.Context)  {
	var user models.User
	id:=ctx.PostForm("id")
	user =models.SearchUserById(id)
	encodeString:=qrcode.PushPic(user)

	ctx.HTML(http.StatusOK,"loginsuccess.tmpl",gin.H{
		"Name":user.Name,
		"Email":user.Email,
		"code":200,
		"File":encodeString,
	})
}
