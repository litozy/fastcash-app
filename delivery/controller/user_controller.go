package controller

import (
	"errors"
	"fmt"
	"net/http"
	"peminjaman/apperror"
	"peminjaman/model"
	"peminjaman/usecase"

	"github.com/gin-gonic/gin"
)

type UserHandler interface {
}

type userHandlerImpl struct {
	usrUsecase usecase.UserUsecase
}

func (usrHandler *userHandlerImpl) GetAllUser(ctx *gin.Context) {
	arrUser, err := usrHandler.usrUsecase.GetAllUser()
	if err != nil {
		fmt.Printf("userHandlerImpl.GetAllUser() : %v", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": "Terjadi kesalahan ketika mengambil data user",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":    arrUser,
		"success": true,
	})
}

func (usrHandler *userHandlerImpl) GetUserByName(ctx *gin.Context) {
	name := ctx.Param("name")
	if name == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Nama tidak boleh kosong",
		})
		return
	}

	usr, err := usrHandler.usrUsecase.GetUserByName(name)
	if err != nil {
		fmt.Printf("UserHandler.GetUserByName() : %v ", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": "Terjadi kesalahan ketika mengambil data User",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    usr,
	})
}

func (usrHandler *userHandlerImpl) RegisterUser(ctx *gin.Context) {
	usr := model.UserModel{}
	err := ctx.ShouldBindJSON(&usr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Invalid JSON data",
		})
		return
	}

	if len(usr.UserName) > 15 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Panjang Nama tidak boleh lebih dari 15 karakter",
		})
		return
	}
	if usr.UserName == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Nama tidak boleh kosong",
		})
		return
	}

	err = usrHandler.usrUsecase.RegisterUser(&usr)
	if err != nil {

		appError := apperror.AppError{}
		if errors.As(err, &appError) {
			fmt.Printf("userHandlerImpl.InsertUser() : %v", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": appError.Error(),
			})
		} else {
			fmt.Printf("userHandlerImpl.AddUser() : %v", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": "Terjadi kesalahan ketika menambahkan data user",
			})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success":        true,
		"successMessage": "Sukses menambahkan data",
	})
}

func NewUserController(srv *gin.Engine, usrUsecase usecase.UserUsecase) UserHandler {
	usrHandler := &userHandlerImpl{
		usrUsecase: usrUsecase,
	}
	srv.POST("/user", usrHandler.RegisterUser)
	srv.GET("/user", usrHandler.GetAllUser)
	srv.GET("/users/:name", usrHandler.GetUserByName)
	return usrHandler
}
