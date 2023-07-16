package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"math/rand"
	"net/http"
	"time"
)

// User 建立结构体与数据表进行映射
type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(20);not null"`
	Telephone string `gorm:"varchar(11);not null;unique"`
	Password  string `gorm:"size:255;not null"`
}

func main() {
	db := InitDb()
	err := db.AutoMigrate(&User{})
	if err != nil {
		fmt.Println("自动创建数据表失败")
	}
	r := gin.Default()
	r.POST("/api/auth/register", func(ctx *gin.Context) {
		// 获取参数
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
			name = RandomString(10)
		}

		log.Println("参数为："+name, telephone, password)
		// 判断手机号是否存在
		if isTelephoneExist(db, telephone) {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"code":    422,
				"message": "手机号已经存在",
			})
			return
		}

		user := User{
			Model:     gorm.Model{},
			Name:      name,
			Telephone: telephone,
			Password:  password,
		}

		db.Create(&user)

		ctx.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": name + " 注册成功",
		})
	})
	r.Run()
}

func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}

func RandomString(n int) string {
	var letters = []byte("abasdajskldajsdljigvhiohASDKLJSAKLDJOIHGIOHFGIOWHIOQ")
	result := make([]byte, n)
	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

func InitDb() *gorm.DB {
	//driverName := "mysql"
	host := "localhost"
	port := "3306"
	database := "OceanLearn"
	username := "root"
	password := "vlt159.0"
	charset := "utf8"
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset)
	db, err := gorm.Open(mysql.New(mysql.Config{DSN: args}), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db

}
