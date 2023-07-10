package controller

import (
	"fmt"
	"net/http"
	"peminjaman/model"
	"peminjaman/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
)

type LateInterestHandler struct {
	lintrsUsecase usecase.LateInterestUsecase
}

func (lintrsHandler LateInterestHandler) GetLateInterestById(ctx *gin.Context) {
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

	lintrs, err := lintrsHandler.lintrsUsecase.GetLateInterestById(id)
	if err != nil {
		fmt.Printf("LateInterestHandler.GetInterestById() : %v ", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": "Terjadi kesalahan ketika mengambil data interest",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    lintrs,
	})
}

func (lintrsHandler LateInterestHandler) GetAllLateInterest(ctx *gin.Context) {
	lintrs, err := lintrsHandler.lintrsUsecase.GetAllLateInterest()
	if err != nil {
		fmt.Printf("lintrsHandler.lintrsUseCase.getAllInterest() : %v", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": "Terjadi kesalahan dalam mengambil semua data interest",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    lintrs,
	})
}

func (lintrsHandler LateInterestHandler) InsertLateInterest(ctx *gin.Context) {
	lintrs := &model.InterestModel{}
	err := ctx.ShouldBindJSON(&lintrs)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Invalid JSON data",
		})
		return
	}
	err = lintrsHandler.lintrsUsecase.InsertLateInterest(lintrs)
	if err != nil {
		fmt.Printf("LateInterestHandler.InsertInterest() : %v ", err.Error())
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

func (lintrsHandler LateInterestHandler) DeleteLateInterest(ctx *gin.Context) {
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

	err = lintrsHandler.lintrsUsecase.DeleteLateInterest(id)
	if err != nil {
		fmt.Printf("lintrsHandler.lintrsUseCase.getAllInterest() : %v", err.Error())
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

func (lintrsHandler LateInterestHandler) UpdateLateInterest(ctx *gin.Context) {
	lintrs := &model.InterestModel{}
	err := ctx.ShouldBindJSON(&lintrs)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Invalid JSON data",
		})
		return
	}

	err = lintrsHandler.lintrsUsecase.UpdateLateInterest(lintrs)
	if err != nil {
		fmt.Printf("lintrsHandler.lintrsUseCase.getAllInterest() : %v", err.Error())
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

func NewLateInterestHandler(srv *gin.Engine, lintrsUsecase usecase.LateInterestUsecase) *LateInterestHandler {
	lintrsHandler := &LateInterestHandler{
		lintrsUsecase: lintrsUsecase,
	}
	srv.GET("/linterest/:id", lintrsHandler.GetLateInterestById)
	srv.GET("/linterest", lintrsHandler.GetAllLateInterest)
	srv.POST("/linterest", lintrsHandler.InsertLateInterest)
	srv.DELETE("/linterest/:id", lintrsHandler.DeleteLateInterest)
	srv.PUT("/linterest", lintrsHandler.UpdateLateInterest)
	return lintrsHandler
}
