package controller

import (
	"OceanLearn/common"
	"OceanLearn/model"
	"OceanLearn/response"
	"OceanLearn/vo"
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
	var requestCategory vo.CreateCategoryRequest
	if err := ctx.ShouldBind(&requestCategory); err != nil {
		response.Fail(ctx, nil, "数据验证错误，分类名称必填")
		return
	}

	category := model.Category{Name: requestCategory.Name}
	// 添加到数据库中
	if err := c.DB.Create(&category).Error; err != nil {
		panic(err)
	}
	response.Success(ctx, gin.H{"category": category}, "")
}

func (c CategoryController) Update(ctx *gin.Context) {
	var requestCategory vo.CreateCategoryRequest
	if err := ctx.ShouldBind(&requestCategory); err != nil {
		response.Fail(ctx, nil, "数据验证错误，分类名称必填")
		return
	}

	// 创建对象，查找数据库并赋值给对象
	var updateCategory model.Category
	categoryID, _ := strconv.Atoi(ctx.Params.ByName("id"))
	result := c.DB.First(&updateCategory, categoryID)

	// 查询错误会返回 ErrRecordNotFound 错误
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		response.Fail(ctx, nil, "分类不存在")
		return
	}

	if err := c.DB.Model(&updateCategory).Update("name", requestCategory.Name).Error; err != nil {
		panic(err)
	}
	response.Success(ctx, gin.H{"category": updateCategory}, "")

}

func (c CategoryController) Show(ctx *gin.Context) {
	var category model.Category
	categoryID, _ := strconv.Atoi(ctx.Params.ByName("id"))
	result := c.DB.First(&category, categoryID)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		response.Fail(ctx, nil, "分类不存在")
		return
	}
	response.Success(ctx, gin.H{"category": category}, "")
}

func (c CategoryController) Delete(ctx *gin.Context) {
	categoryID, _ := strconv.Atoi(ctx.Params.ByName("id"))
	// 传入模型对象，告诉GORM要删除的是哪张表
	// 根据主键删除，指定表格和主键值，DELETE FROM users WHERE id = 10;
	if err := c.DB.Delete(model.Category{}, categoryID).Error; err != nil {
		response.Fail(ctx, nil, "删除失败")
		return
	}
	response.Success(ctx, nil, "删除成功")
}
