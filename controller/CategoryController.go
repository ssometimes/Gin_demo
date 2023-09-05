package controller

import (
	"OceanLearn/common"
	"OceanLearn/model"
	"OceanLearn/response"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

type ICategoryController interface {
	RestController
}

type CategoryController struct {
	DB *gorm.DB
}

// NewCategoryController 这里返回的类型是 ICategoryController？
func NewCategoryController() ICategoryController {
	db := common.GetDb()
	err := db.AutoMigrate(model.Category{})
	if err != nil {
		return nil
	}
	return CategoryController{DB: db}
}

func (c CategoryController) Create(ctx *gin.Context) {
	// 定义参数接口请求的参数值
	var requestCategory model.Category
	ctx.Bind(requestCategory)
	if requestCategory.Name == "" {
		response.Fail(ctx, nil, "数据验证错误，分类名称必填")
	}
	// 添加到数据库中
	c.DB.Create(requestCategory)
	response.Success(ctx, gin.H{"category": requestCategory}, "")
}

func (c CategoryController) Update(ctx *gin.Context) {
	var requestCategory model.Category
	ctx.ShouldBind(requestCategory)
	if requestCategory.Name == "" {
		response.Fail(ctx, nil, "数据验证错误，分类名称必填")
	}

	// 创建对象，查找数据库并赋值给对象
	var updateCategory model.Category
	categoryID, _ := strconv.Atoi(ctx.Params.ByName("id"))
	result := c.DB.First(&updateCategory, categoryID)

	// 查询错误会返回 ErrRecordNotFound 错误
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		response.Fail(ctx, nil, "分类不存在")
	}

	c.DB.Model(&updateCategory).Updates(gin.H{"name": requestCategory.Name})
	response.Success(ctx, gin.H{"category": updateCategory}, "")

}

func (c CategoryController) Show(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (c CategoryController) Delete(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}
