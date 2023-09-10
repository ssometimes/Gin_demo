package vo

// 为了避免多个参数需要校验时，重复写过多的代码，使用一个结构体统一绑定验证
type CreateCategoryRequest struct {
	Name string `json:"name" binding:"required"`
}
