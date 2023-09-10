package main

import (
	"OceanLearn/controller"
	"OceanLearn/middleware"
	"github.com/gin-gonic/gin"
)

// POST此类函数是将api添加到一个组里面，所以调用后返回这个 gin.Engine就获得了所有方法
func CollectRoute(r *gin.Engine) *gin.Engine {
	r.Use(middleware.CORSMiddleware(), middleware.RecoveryMiddleware())
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", controller.Login)
	r.GET("/api/auth/info", middleware.AuthMiddleware(), controller.Info)

	// 建立路由组
	categoryRoutes := r.Group("/categories")
	categoryController := controller.NewCategoryController()
	categoryRoutes.POST("", categoryController.Create)
	categoryRoutes.GET("/:id", categoryController.Show)
	categoryRoutes.PUT("/:id", categoryController.Update)
	categoryRoutes.DELETE("/:id", categoryController.Delete)
	return r
}
