package controller

import (
	"fmt"
	"net/http"
	"peminjaman/model"
	"peminjaman/usecase"

	"github.com/gin-gonic/gin"
)

type LoanProductHandler struct {
	lprdctUsecase usecase.LoanProductUsecase
}

func (lprdctHandler LoanProductHandler) GetLoanProductById(ctx *gin.Context) {
	id := &model.LoanProductModel{}
	err := ctx.ShouldBindJSON(&id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Invalid JSON data",
		})
		return
	}
	if id == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Id tidak boleh kosong",
		})
		return
	}
	// idText := ctx.Param("id")
	// id, err := strconv.Atoi(idText)
	// if err != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{
	// 		"success":      false,
	// 		"errorMessage": "Id harus angka",
	// 	})
	// 	return
	// }

	lprdct, err := lprdctHandler.lprdctUsecase.GetLoanProductById(id.Id)
	if err != nil {
		fmt.Printf("LoanProductHandler.GetLoanProductById() : %v ", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": "Terjadi kesalahan ketika mengambil data loanProduct",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    lprdct,
	})
}

func (lprdctHandler LoanProductHandler) GetAllLoanProduct(ctx *gin.Context) {
	lprdct, err := lprdctHandler.lprdctUsecase.GetAllLoanProduct()
	if err != nil {
		fmt.Printf("lprdctHandler.lprdctUseCase.getAllLoanProduct() : %v", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": "Terjadi kesalahan dalam mengambil semua data loanProduct",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    lprdct,
	})
}

func (lprdctHandler LoanProductHandler) InsertLoanProduct(ctx *gin.Context) {
	lprdct := &model.LoanProductModel{}
	err := ctx.ShouldBindJSON(&lprdct)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Invalid JSON data",
		})
		return
	}
	err = lprdctHandler.lprdctUsecase.InsertLoanProduct(lprdct)
	if err != nil {
		fmt.Printf("LoanProductHandler.InsertLoanProduct() : %v ", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": "Terjadi kesalahan ketika menyimpan data loanProduct",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func (lprdctHandler LoanProductHandler) DeleteLoanProduct(ctx *gin.Context) {
	id := &model.LoanProductModel{}
	err := ctx.ShouldBindJSON(&id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Invalid JSON data",
		})
		return
	}
	if id == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Id tidak boleh kosong",
		})
		return
	}

	// idText := ctx.Param("id")
	// if idText == "" {
	// 	ctx.JSON(http.StatusBadGateway, gin.H{
	// 		"success":      false,
	// 		"errorMessage": "Id tidak boleh kosong",
	// 	})
	// 	return
	// }

	// id, err := strconv.Atoi(idText)
	// if err != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{
	// 		"succes":       false,
	// 		"errorMessage": "Id harus angka",
	// 	})
	// 	return
	// }

	err = lprdctHandler.lprdctUsecase.DeleteLoanProduct(id.Id)
	if err != nil {
		fmt.Printf("lprdctHandler.lprdctUseCase.getAllLoanProduct() : %v", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": "Terjadi kesalahan dalam menghapus data loanProduct",
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func (lprdctHandler LoanProductHandler) UpdateLoanProduct(ctx *gin.Context) {
	lprdct := &model.LoanProductModel{}
	err := ctx.ShouldBindJSON(&lprdct)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Invalid JSON data",
		})
		return
	}

	err = lprdctHandler.lprdctUsecase.UpdateLoanProduct(lprdct)
	if err != nil {
		fmt.Printf("lprdctHandler.lprdctUseCase.getAllLoanProduct() : %v", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": "Terjadi kesalahan dalam memperbarui data loanProduct",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func NewLoanProductHandler(srv *gin.Engine, lprdctUsecase usecase.LoanProductUsecase) *LoanProductHandler {
	lprdctHandler := &LoanProductHandler{
		lprdctUsecase: lprdctUsecase,
	}
	srv.GET("/loanProduct", lprdctHandler.GetLoanProductById)
	srv.GET("/loanProduct", lprdctHandler.GetAllLoanProduct)
	srv.POST("/loanProduct", lprdctHandler.InsertLoanProduct)
	srv.DELETE("/loanProduct", lprdctHandler.DeleteLoanProduct)
	srv.PUT("/loanProduct", lprdctHandler.UpdateLoanProduct)
	return lprdctHandler
}
