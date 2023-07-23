package main

import (
	"OceanLearn/controller"
	"github.com/gin-gonic/gin"
)

// POST此类函数是将api添加到一个组里面，所以调用后返回这个 gin.Engine就获得了所有方法
func CollectRoute(r *gin.Engine) *gin.Engine {
	r.POST("/api/auth/register", controller.Register)
	return r
}
