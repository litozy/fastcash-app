package controller

import (
	"fmt"
	"net/http"
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
	ta := model.TransactionApply{}
	err := ctx.ShouldBindJSON(&ta)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"errorMessage": "Invalid JSON data",
		})
		return
	}

	err = taController.taUsecase.InsertApplication(&ta)
	if err != nil {
		fmt.Printf("serviceHandlerImpl.AddService() : %v", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"errorMessage": "Terjadi kesalahan ketika menambahkan data service",
			"error": err.Error(),
		})
		
	return
	}
}

func NewTransactionApplyController(srv *gin.Engine, taUsecase usecase.TransactionApplyUsecase) TransactionApplyController {
	taController := &transactionApplyControllerImpl{
		taUsecase: taUsecase,
	}

	srv.POST("/application", taController.InsertApplication)
	
	return taController
}