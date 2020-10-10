package middleware

import (
	"github.com/gin-gonic/gin"
)

 func JwtVerify(ctx *gin.Context)  {
		token:=ctx.GetHeader("Authorization")
		if token == "" {
			//ctx.Redirect(http.StatusFound, "/login")
			return
		}
		ctx.Header("Authorization",token)
		ctx.Next()
	}

