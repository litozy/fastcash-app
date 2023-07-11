package controller

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"peminjaman/apperror"
	"peminjaman/model"
	"peminjaman/usecase"
	"peminjaman/utils/authutil"

	"github.com/gin-gonic/gin"
)

type loginControllerImpl struct {
	lgUsecase usecase.LoginUsecase
}

func (lgController loginControllerImpl) Login(ctx *gin.Context) {
	loginData := &model.LoginModel{}
	err := ctx.ShouldBindJSON(&loginData)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"errorMessage" : "Invalid JSON data",
		})
		return
	}
	
	err = lgController.lgUsecase.GetUserByNameAndPassword(loginData.Username, loginData.Password)
	if err != nil {

		appError := apperror.AppError{}
		if errors.As(err, &appError) {
			fmt.Printf("loginHandlerImpl.Login() : %v", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"errorMessage": appError.ErrorMessage,
			})
		} else {
			fmt.Printf("loginHandlerImpl.Login() : %v", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"errorMessage": "Terjadi kesalahan ketika login",
			})
		}
	return
	}

	temp, err := authutil.GenerateToken(loginData.Username)
	if err != nil {
		log.Println("Token Invalid")
	}
	ctx.JSON(http.StatusOK, gin.H{
		"token" : temp,
	})
}

func NewLoginController(srv *gin.Engine, lgUsecase usecase.LoginUsecase) {
	lgHandler := &loginControllerImpl{
		lgUsecase: lgUsecase,
	}
	srv.POST("/login", lgHandler.Login)
}