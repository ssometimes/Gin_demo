package middleware

import (
	"OceanLearn/common"
	"OceanLearn/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// 中间件的标准写法，应该通过闭包实现，这样能在返回函数前做操作
func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取authorization header
		tokenString := ctx.GetHeader("Authorization")

		// token为空或者请求头字段的开头不是Bearer
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer") {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code":    401,
				"message": "权限不足",
			})
			ctx.Abort()
			return
		}

		tokenString = tokenString[7:]
		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code":    401,
				"message": "权限不足",
			})
			ctx.Abort()
			return
		}

		// 验证通过后获取claim 中的userId, 并去数据库中查找
		userId := claims.UserId
		DB := common.GetDb()
		var user model.User
		DB.First(&user, userId)

		// 用户存在情况
		if user.ID == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code":    401,
				"message": "权限不足",
			})
			ctx.Abort()
			return
		}

		// 用户存在，将user的信息写入上下文
		ctx.Set("user", user)
		ctx.Next()
	}
}
