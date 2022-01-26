package model

import (
	"time"

	"gorm.io/gorm"
)

// User 用户
type User struct {
	ID        uint   `gorm:"primaryKey;autoIncrement;" json:"id"`
	Name      string `gorm:"type:varchar(10);comment:名称" json:"name" binding:"required"`
	Nickname  string `gorm:"type:varchar(20);comment:昵称" json:"nickname" binding:"required,nefield=Name"`
	Password  string `gorm:"type:varchar(32);comment:密码" json:"password" binding:"required"`
	Birthday  string `gorm:"type:date;comment:出生日期" json:"birthday" binding:"omitempty,birthday"`
	Gender    int    `gorm:"type:tinyint(1);comment:0女 1男" json:"gender" binding:"omitempty,min=0,max=1"`
	Phone     string `gorm:"type:char(13);uniqueIndex;comment:手机号" json:"phone"`
	Email     string `gorm:"type:varchar(30);uniqueIndex;comment:邮箱" json:"email" binding:"omitempty,email"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Users struct {
	List []User `json:"list" binding:"gt=0,lt=100,size,dive"` // 切片长度>0，<100,dive验证每一项
	Size int    `json:"size"`
}

func NewUser(ID uint, name string) *User {
	return &User{ID: ID, Name: name}
}

type UserQuery struct {
	Name     string `json:"name" form:"name"`
	Page     int    `json:"page" form:"page" binding:"required"`
	PageSize int    `json:"page_size" form:"page_size" binding:"required"`
}
