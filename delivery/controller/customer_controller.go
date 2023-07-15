package controller

import (
	"fmt"
	"net/http"
	"peminjaman/model"
	"peminjaman/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CustomerHandler struct {
	cstmUsecase usecase.CustomerUsecase
}

func (cstmHandler CustomerHandler) GetCustomerById(ctx *gin.Context) {
	idText := ctx.Param("id")
	if idText == "" {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"success":      false,
			"errorMessage": "Id tidak boleh kosong",
		})
		return
	}

	id, err := strconv.Atoi(idText)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"succes":       false,
			"errorMessage": "Id harus angka",
		})
		return
	}

	cstm, err := cstmHandler.cstmUsecase.GetCustomerById(id)
	if err != nil {
		fmt.Printf("CustomerHandler.GetCustomerById() : %v ", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": "Terjadi kesalahan ketika mengambil data customer",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    cstm,
	})
}

func (cstmHandler CustomerHandler) GetAllCustomer(ctx *gin.Context) {
	cstm, err := cstmHandler.cstmUsecase.GetAllCustomer()
	if err != nil {
		fmt.Printf("cstmHandler.cstmUseCase.getAllCustomer() : %v", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": "Terjadi kesalahan dalam mengambil semua data customer",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    cstm,
	})
}

func (cstmHandler CustomerHandler) InsertCustomer(ctx *gin.Context) {
	cstm := &model.CustomerModel{}
	err := ctx.ShouldBindJSON(&cstm)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Invalid JSON data",
			"error" : err.Error(),
		})
		return
	}
	err = cstmHandler.cstmUsecase.InsertCustomer(cstm)
	if err != nil {
		fmt.Printf("CustomerHandler.InsertCustomer() : %v ", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": "Terjadi kesalahan ketika menyimpan data customer",
			"error" : err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func (cstmHandler CustomerHandler) DeleteCustomer(ctx *gin.Context) {
	idText := ctx.Param("id")
	if idText == "" {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"success":      false,
			"errorMessage": "Id tidak boleh kosong",
		})
		return
	}

	id, err := strconv.Atoi(idText)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"succes":       false,
			"errorMessage": "Id harus angka",
		})
		return
	}

	err = cstmHandler.cstmUsecase.DeleteCustomer(id)
	if err != nil {
		fmt.Printf("cstmHandler.cstmUseCase.getAllCustomer() : %v", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": "Terjadi kesalahan dalam menghapus data customer",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func (cstmHandler CustomerHandler) UpdateCustomer(ctx *gin.Context) {
	cstm := &model.CustomerModel{}
	err := ctx.ShouldBindJSON(&cstm)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Invalid JSON data",
		})
		return
	}

	err = cstmHandler.cstmUsecase.UpdateCustomer(cstm)
	if err != nil {
		fmt.Printf("cstmHandler.cstmUseCase.getAllCustomer() : %v", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": "Terjadi kesalahan dalam memperbarui data customer",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func (cstmHandler CustomerHandler) UpdateCustomerStatus(ctx *gin.Context) {
	cstm := &model.CustomerModel{}
	err := ctx.ShouldBindJSON(&cstm)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Invalid JSON data",
		})
		return
	}

	err = cstmHandler.cstmUsecase.UpdateCustomerStatus(cstm)
	if err != nil {
		fmt.Printf("cstmHandler.cstmUseCase.getAllCustomer() : %v", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": "Terjadi kesalahan dalam memperbarui data customer status",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func NewCustomerHandler(srv *gin.Engine, cstmUsecase usecase.CustomerUsecase) *CustomerHandler {
	cstmHandler := &CustomerHandler{
		cstmUsecase: cstmUsecase,
	}
	srv.GET("/customer/:id", cstmHandler.GetCustomerById)
	srv.GET("/customer", cstmHandler.GetAllCustomer)
	srv.POST("/customer", cstmHandler.InsertCustomer)
	srv.DELETE("/customer/:id", cstmHandler.DeleteCustomer)
	srv.PUT("/customer", cstmHandler.UpdateCustomer)
	srv.PUT("/customerStatus", cstmHandler.UpdateCustomerStatus)
	return cstmHandler
}
