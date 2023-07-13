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

type TransactionPaymentController interface {
}

type transactionPaymentControllerImpl struct {
	trpUsecase usecase.TransactionPaymentUsecase
}

func (trpController *transactionPaymentControllerImpl) InsertPayment(ctx *gin.Context) {
	trp := &model.TransactionPayment{}
	err := ctx.ShouldBindJSON(&trp)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"errorMessage": "Invalid JSON data",
		})
		return
	}

	err = trpController.trpUsecase.InsertPayment(trp)
	if err != nil {
		appError := apperror.AppError{}
		if errors.As(err, &appError) {
			fmt.Printf("transactionPaymentControllerImpl.InsertPayment() : %v", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"errorMessage": appError.ErrorMessage,
			})
		
		} else {
			fmt.Printf("transactionPaymentControllerImpl.InsertPayment() : %v", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"errorMessage": "Terjadi kesalahan ketika menambahkan data transaksi payment",
				"error" : err.Error(),
			})
		
		}
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"successMessage": "Success memasukkan data",
		"success" : true,
		
	})
}

func NewTransactionPaymentController(srv *gin.Engine, trpUsecase usecase.TransactionPaymentUsecase) TransactionPaymentController {
	trpController := &transactionPaymentControllerImpl{
		trpUsecase: trpUsecase,
	}

	srv.POST("/payment", trpController.InsertPayment)

	return trpController
}