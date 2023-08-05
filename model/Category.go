package model

import "time"

type Category struct {
	// GORM 使用 ID 作为主键
	ID   uint   `json:"id" gorm:"primary_key"`
	Name string `json:"name" gorm:"type:varchar(50);not null;unique"`
	// GORM 约定使用CreatedAt、UpdateAt 字段追踪创建、更新时间
	// 使用规定的情况下，数据库有创建数据或者修改数据，GORM 会自动帮我记录时间
	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp"`
	UpdateAt  time.Time `json:"update_at" gorm:"type:timestamp"`
}
