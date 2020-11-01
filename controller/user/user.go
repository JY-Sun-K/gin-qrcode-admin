package user

import (
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"log"
	"net/http"
	"qcode/models"
	"qcode/pkg/jwt"
	"qcode/pkg/qrcode"
)

func SigninGet(ctx *gin.Context)  {
	ctx.HTML(http.StatusOK,"signin.tmpl",gin.H{
		"code":200,
	})
}

func SiginPost(ctx *gin.Context)  {
	//db:=models.GetDB()

	name :=ctx.PostForm("name")
	password := ctx.PostForm("password")
	email := ctx.PostForm("email")
	phone := ctx.PostForm("phone")
	adderss := ctx.PostForm("address")

	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,gin.H{
			"msg":"加密错误",
		})
		return
	}
	models.CreateUser(name,string(hasedPassword),email,phone,adderss)
	//newuser := &models.User{
	//	Name:     name,
	//	Password: string(hasedPassword),
	//	Email:    email,
	//	Phone:    phone,
	//	Adderss:  adderss,
	//}
	//db.Create(&newuser)
	//qdata := name
	qdata := "http://127.0.0.1:8080/user/loginsusscess/"+name
	err = qrcode.CreateQrCode(qdata,name,200,200)
	if err != nil {
		fmt.Println("File reading error", err)
	}
	// 重定向，网页跳转
	ctx.Redirect(http.StatusFound, "/user/login")





	
	
	
}
func LoginSuccess(ctx *gin.Context)  {
	var user models.User

	db:=models.GetDB()


	userId,_ :=ctx.Get("userId")

	db.Where("ID = ?", userId).First(&user)

	fmt.Println(userId)
	PicName := user.Name+".png"
	fileName := "./qrcodepng/"+PicName
	PicFile, err := ioutil.ReadFile(fileName) // just pass the file name
	if err != nil {
		fmt.Print(err)
	}
	encodeString := base64.StdEncoding.EncodeToString(PicFile)

	ctx.HTML(http.StatusOK,"loginsuccess.tmpl",gin.H{
		"Name":user.Name,
		"Email":user.Email,
		"code":200,
		"File":encodeString,
	})
}

func LoginGet(ctx *gin.Context) {
	ctx.HTML(http.StatusOK,"login.tmpl",gin.H{
		"code":200,
	})
}


func LoginPost(ctx *gin.Context) {
	var user models.User
	db:=models.GetDB()

	name :=ctx.PostForm("name")
	password := ctx.PostForm("password")

	db.Where("name = ? AND root = false", name).First(&user)
	if user.ID == 0 {
		ctx.JSON(http.StatusNotFound,gin.H{
			"code":404,
			"msg":"用户不存在",
		})
	}


	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); user.Name==name &&err!=nil {
		ctx.JSON(http.StatusUnprocessableEntity,gin.H{
			"code":442,
			"msg":"登录失败",
		})
		ctx.Abort()
		return

	}

	token ,err:= jwt.ReleaseToken(user)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError,gin.H{
			"code":500,
			"msg":"token发放错误",
		})
		ctx.Abort()
		return
	}

	ctx.SetCookie("user_cookie",token,25*60,"/","127.0.0.1",false,true)
	cookie ,err:=ctx.Request.Cookie("user_cookie")
	if err != nil {
		log.Println(err)
	}
	log.Println(cookie)
	log.Print(token)
	//ctx.Request.Header.Add("Authorization",token)
	//log.Println(ctx.Request.Header)
	//ctx.Next()
	// 重定向，网页跳转
	ctx.Redirect(http.StatusMovedPermanently, "/user/loginsusscess")

	//ctx.Request.URL.Path = "/loginsusscess"

	//ctx.JSON(http.StatusOK,gin.H{
	//	"code":200,
	//	"msg":"登录成功",
	//	"token":token,
	//})


}


