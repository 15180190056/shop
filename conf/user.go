package conf

import (
	"gorm.io/gorm"
)

// User 是数据库中的用户模型
type User struct {
	gorm.Model
	Orders   []Order `gorm:"foreignkey:UserID"`
	Username string  `gorm:"unique"`
	Password string
	Email    string
	Phone    string `gorm:"unique"`
}
