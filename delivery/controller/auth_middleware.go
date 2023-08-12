package controller

import (
	"fmt"
	"net/http"
	"peminjaman/utils/authutil"
	"strings"

	"github.com/gin-gonic/gin"
)

type authHeader struct {
	Authorization string `header:"Authorization"`
}

func RequireToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		h := authHeader{}
		if err := ctx.ShouldBindHeader(&h); err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message" : "Unauthorize",
			})
			ctx.Abort()
			return
		}
		tokenString := strings.Replace(h.Authorization, "Bearer " , "", -1)
		fmt.Println(tokenString)
		if tokenString == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message" : "Unauthorize",
			})
			ctx.Abort()
			return
		}
		token, err := authutil.VerifyAccessToken(tokenString)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message" : "Unauthorize",
			})
			ctx.Abort()
			return
		}
		fmt.Println(token)
		if token != "" {
			ctx.Next()
		} else {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message" : "Unauthorize 3",
			})
			ctx.Abort()
			return
		}
	}
}