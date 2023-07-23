package main

import (
	"OceanLearn/common"
	"OceanLearn/model"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "gorm.io/driver/mysql"
)

func main() {
	// 初始化数据库配置，并自动对结构体与数据表进行映射
	db := common.GetDb()
	err := db.AutoMigrate(&model.User{})
	if err != nil {
		fmt.Println("自动创建数据表失败")
	}
	r := gin.Default()
	r = CollectRoute(r)
	panic(r.Run())
}
