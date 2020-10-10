package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"qcode/models"
	"qcode/pkg/jwt"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.Println(ctx.Request.Header)
		tokenstring := ctx.Request.Header.Get("Authorization")

		if tokenstring == "" {
			ctx.JSON(http.StatusUnauthorized,gin.H{
				"code":401,
				"msg":"权限不住",
			})
			ctx.Abort()
			return
		}

		tokenstring = tokenstring[7:]
		token ,claims,err :=jwt.ParseToken(tokenstring)
		if err !=nil ||!token.Valid {
			ctx.JSON(http.StatusUnauthorized,gin.H{
				"code":401,
				"msg":"权限不住",
			})
			ctx.Abort()
			return
		}
		userId :=claims.UserId

		DB := models.GetDB()
		var user models.User
		DB.First(&user,userId)
		if user.ID==0 {
			ctx.JSON(http.StatusUnauthorized,gin.H{
				"code":401,
				"msg":"用户不存在，权限不住",
			})
			ctx.Abort()
			return
		}
		ctx.Set("userId",user.ID)
		ctx.Next()
	}
}
