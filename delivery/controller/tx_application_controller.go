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

type TransactionApplyController interface {

}

type transactionApplyControllerImpl struct {
	taUsecase usecase.TransactionApplyUsecase
}

func (taController *transactionApplyControllerImpl) InsertApplication(ctx *gin.Context) {
	tra := &model.TransactionApply{}
	err := ctx.ShouldBindJSON(&tra)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"errorMessage": "Invalid JSON data",
		})
		return
	}

	err = taController.taUsecase.InsertApplication(tra)
	if err != nil {
		appError := apperror.AppError{}
		if errors.As(err, &appError) {
			fmt.Printf("transactionApplyControllerImpl.InsertApplication() : %v", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"errorMessage": appError.ErrorMessage,
		})
		return
	} else {
		fmt.Printf("serviceHandlerImpl.AddService() : %v", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"errorMessage": "Terjadi kesalahan ketika menambahkan data service",
			"error": err.Error(),
		})
	}
	return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"successMessage": "Success memasukkan data",
		"success" : true,
		
	})
}

func NewTransactionApplyController(srv *gin.Engine, taUsecase usecase.TransactionApplyUsecase) TransactionApplyController {
	taController := &transactionApplyControllerImpl{
		taUsecase: taUsecase,
	}

	srv.POST("/application", taController.InsertApplication)
	
	return taController
}