package model

import "gorm.io/gorm"

// User 建立结构体与数据表进行映射
type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(20);not null"`
	Telephone string `gorm:"varchar(11);not null;unique"`
	Password  string `gorm:"size:255;not null"`
}
