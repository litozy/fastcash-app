package controller

import (
	"fmt"
	"net/http"
	"peminjaman/model"
	"peminjaman/usecase"

	"github.com/gin-gonic/gin"
)

type OjkStatusHandler struct {
	ojkstatUsecase usecase.OjkStatusUsecase
}

func (ojkstatHandler OjkStatusHandler) GetOjkStatusById(ctx *gin.Context) {
	id := &model.OjkStatusModel{}
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

	ojkstat, err := ojkstatHandler.ojkstatUsecase.GetOjkStatusById(id.Id)
	if err != nil {
		fmt.Printf("OjkStatusHandler.GetOjkStatusById() : %v ", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": "Terjadi kesalahan ketika mengambil data ojkStatus",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    ojkstat,
	})
}

func (ojkstatHandler OjkStatusHandler) GetAllOjkStatus(ctx *gin.Context) {
	ojkstat, err := ojkstatHandler.ojkstatUsecase.GetAllOjkStatus()
	if err != nil {
		fmt.Printf("ojkstatHandler.ojkstatUseCase.getAllOjkStatus() : %v", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": "Terjadi kesalahan dalam mengambil semua data ojkStatus",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    ojkstat,
	})
}

func (ojkstatHandler OjkStatusHandler) InsertOjkStatus(ctx *gin.Context) {
	ojkstat := &model.OjkStatusModel{}
	err := ctx.ShouldBindJSON(&ojkstat)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Invalid JSON data",
		})
		return
	}
	err = ojkstatHandler.ojkstatUsecase.InsertOjkStatus(ojkstat)
	if err != nil {
		fmt.Printf("OjkStatusHandler.InsertOjkStatus() : %v ", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": "Terjadi kesalahan ketika menyimpan data ojkStatus",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func (ojkstatHandler OjkStatusHandler) DeleteOjkStatus(ctx *gin.Context) {
	id := &model.OjkStatusModel{}
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

func NewOjkStatusHandler(srv *gin.Engine, ojkstatUsecase usecase.OjkStatusUsecase) *OjkStatusHandler {
	ojkstatHandler := &OjkStatusHandler{
		ojkstatUsecase: ojkstatUsecase,
	}
	srv.GET("/ojkStatus", ojkstatHandler.GetOjkStatusById)
	srv.GET("/ojkStatus", ojkstatHandler.GetAllOjkStatus)
	srv.POST("/ojkStatus", ojkstatHandler.InsertOjkStatus)
	return ojkstatHandler
}
