// Package model provides ...
package model

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type UserModel struct {
	//*BaseModel
	Id        int64     `gorm:"column:id" json:"id"`
	UserName  string    `gorm:"column:username" json:"username"`
	Password  string    `gorm:"colunmn:password" json:"password"`
	Name      string    `gorm:"colunmn:name" json:"name"`
	Avatar    string    `gorm:"colunmn:avatar" json:"avatar"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (user *UserModel) TableName() string {
	return "users"
}

func (u *UserModel) BeforeUpdate(*gorm.DB) error {
	fmt.Println("[" + u.UpdatedAt.String() + "]触发更新事件")
	return nil
}
