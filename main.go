package main

import (
	"OceanLearn/common"
	"OceanLearn/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	_ "gorm.io/driver/mysql"
	"os"
)

func main() {
	// 项目开始前进行初始化配置
	InitConfig()
	// 初始化数据库配置，并自动对结构体与数据表进行映射
	db := common.GetDb()
	err := db.AutoMigrate(&model.User{})
	if err != nil {
		fmt.Println("自动创建数据表失败")
	}
	r := gin.Default()
	r = CollectRoute(r)
	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	}
	// panic(r.Run())
}

func InitConfig() {
	workdir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(workdir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic("读取配置失败")
	}
}
