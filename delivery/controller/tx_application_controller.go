package controller

import (
	"errors"
	"fmt"
	"net/http"
	"peminjaman/apperror"
	"peminjaman/model"
	"peminjaman/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TransactionApplyController interface {

}

type transactionApplyControllerImpl struct {
	taUsecase usecase.TransactionApplyUsecase
}

func (taController *transactionApplyControllerImpl) GetAppById(ctx *gin.Context) {
	idText := ctx.Param("id")
	if idText == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"errorMessage": "Id tidak boleh kosong",
		})
		return
	}

	id, err := strconv.Atoi(idText)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"errorMessage": "Id harus berupa angka",
		})
		return
	}

	tra, err := taController.taUsecase.GetTransactionApplyById(id)
	if err != nil {
		appError := apperror.AppError{}
		if errors.As(err, &appError) {
			fmt.Printf("transactionApplyControllerImpl.GetAppById() : %v", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"errorMessage": appError.ErrorMessage,
			})
			return
		} else {
			fmt.Printf("transactionApplyControllerImpl.GetAppById() : %v", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"errorMessage": "Terjadi kesalahan ketika mengambil data transaksi",
			})
		return
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data" : tra,
		"success": true,
	})

}

func (taController *transactionApplyControllerImpl) GetAllApp(ctx *gin.Context) {
	arrTr, err := taController.taUsecase.GetAllApp()
	if err != nil {
		fmt.Printf("serviceHandlerImpl.GetAllService() : %v", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"errorMessage": "Terjadi kesalahan ketika mengambil data transaksi application",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data" : arrTr,
		"success": true,
	})
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
		
		} else {
			fmt.Printf("serviceHandlerImpl.InsertPayment() : %v", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"errorMessage": "Terjadi kesalahan ketika menambahkan data transaksi application",
			})
		
		}
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"successMessage": "Success memasukkan data",
		"success" : true,
		
	})
}

func (taController *transactionApplyControllerImpl) UpdateStatusOjk(ctx *gin.Context) {
	tra := &model.TransactionApply{}
	err := ctx.ShouldBindJSON(&tra)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"errorMessage": "Invalid JSON data",
		})
		return
	}

	err = taController.taUsecase.UpdateStatusOjk(tra)
	if err != nil {
		appError := apperror.AppError{}
		if errors.As(err, &appError) {
			fmt.Printf("transactionApplyControllerImpl.UpdateStatusOjk() : %v", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"errorMessage": appError.ErrorMessage,
			})
		} else {
			fmt.Printf("transactionApplyControllerImpl.UpdateStatusOjk() : %v", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"errorMessage": "Terjadi kesalahan ketika mengubah data transaksi",
			})
		
		}
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"successMessage": "Success mengubah data",
		"success" : true,
		
	})
}

func NewTransactionApplyController(srv *gin.Engine, taUsecase usecase.TransactionApplyUsecase) TransactionApplyController {
	taController := &transactionApplyControllerImpl{
		taUsecase: taUsecase,
	}

	srv.POST("/application", taController.InsertApplication)
	srv.GET("/application", taController.GetAllApp)
	srv.GET("/application/:id", taController.GetAppById)
	srv.PUT("/statusUpdate", taController.UpdateStatusOjk)
	
	return taController
}