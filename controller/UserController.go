package controller

import (
	"OceanLearn/common"
	"OceanLearn/model"
	"OceanLearn/util"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func Register(ctx *gin.Context) {
	// 获取参数
	DB := common.GetDb()
	name := ctx.PostForm("name")
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")
	// 数据验证
	if len(telephone) != 11 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "手机号必须为 11 位",
		})
		return
	}

	if len(password) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "密码不能少于 6 位",
		})
		return
	}

	if len(name) == 0 {
		name = util.RandomString(10)
	}

	log.Println("参数为："+name, telephone, password)
	// 判断手机号是否存在
	if isTelephoneExist(DB, telephone) {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "手机号已经存在",
		})
		return
	}

	// 创建新的对象进行存储
	user := model.User{
		Model:     gorm.Model{},
		Name:      name,
		Telephone: telephone,
		Password:  password,
	}

	DB.Create(&user)

	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": name + " 注册成功",
	})
}

func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
