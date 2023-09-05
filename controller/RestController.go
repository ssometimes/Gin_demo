package controller

import "github.com/gin-gonic/gin"

// RestController 为了给其他功能复用增删改查
type RestController interface {
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Show(ctx *gin.Context)
	Delete(ctx *gin.Context)
}
