package controller

import (
	"fmt"
	"net/http"
	"peminjaman/model"
	"peminjaman/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
)

type InterestHandler struct {
	intrsUsecase usecase.InterestUsecase
}

func (intrsHandler InterestHandler) GetInterestById(ctx *gin.Context) {
	idText := ctx.Param("id")
	if idText == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Id tidak boleh kosong",
		})
		return
	}

	id, err := strconv.Atoi(idText)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Id harus angka",
		})
		return
	}

	intrs, err := intrsHandler.intrsUsecase.GetInterestById(id)
	if err != nil {
		fmt.Printf("InterestHandler.GetInterestById() : %v ", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": "Terjadi kesalahan ketika mengambil data interest",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    intrs,
	})
}

func (intrsHandler InterestHandler) GetAllInterest(ctx *gin.Context) {
	intrs, err := intrsHandler.intrsUsecase.GetAllInterest()
	if err != nil {
		fmt.Printf("intrsHandler.intrsUseCase.getAllInterest() : %v", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": "Terjadi kesalahan dalam mengambil semua data interest",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    intrs,
	})
}

func (intrsHandler InterestHandler) InsertInterest(ctx *gin.Context) {
	intrs := &model.InterestModel{}
	err := ctx.ShouldBindJSON(&intrs)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Invalid JSON data",
		})
		return
	}
	err = intrsHandler.intrsUsecase.InsertInterest(intrs)
	if err != nil {
		fmt.Printf("InterestHandler.InsertInterest() : %v ", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": "Terjadi kesalahan ketika menyimpan data interest",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func (intrsHandler InterestHandler) DeleteInterest(ctx *gin.Context) {
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

	err = intrsHandler.intrsUsecase.DeleteInterest(id)
	if err != nil {
		fmt.Printf("intrsHandler.intrsUseCase.getAllInterest() : %v", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": "Terjadi kesalahan dalam menghapus data interest",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func (intrsHandler InterestHandler) UpdateInterest(ctx *gin.Context) {
	intrs := &model.InterestModel{}
	err := ctx.ShouldBindJSON(&intrs)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Invalid JSON data",
		})
		return
	}

	err = intrsHandler.intrsUsecase.UpdateInterest(intrs)
	if err != nil {
		fmt.Printf("intrsHandler.intrsUseCase.getAllInterest() : %v", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": "Terjadi kesalahan dalam memperbarui data interest",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func NewInterestHandler(srv *gin.Engine, intrsUsecase usecase.InterestUsecase) *InterestHandler {
	intrsHandler := &InterestHandler{
		intrsUsecase: intrsUsecase,
	}
	srv.GET("/interest/:id", intrsHandler.GetInterestById)
	srv.GET("/interest", intrsHandler.GetAllInterest)
	srv.POST("/interest", intrsHandler.InsertInterest)
	srv.DELETE("/interest/:id", intrsHandler.DeleteInterest)
	srv.PUT("/interest", intrsHandler.UpdateInterest)
	return intrsHandler
}
