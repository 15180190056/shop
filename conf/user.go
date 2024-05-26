package conf

import (
	"gorm.io/gorm"
)

// User 是数据库中的用户模型
type User struct {
	gorm.Model
	Username string `gorm:"unique"`
	Password string
	Email    string
	Phone    string `gorm:"unique"`
}

// Category 对应数据库中的categories表
//type Category struct {
//	gorm.Model
//	Name     string    `gorm:"column:name"`
//	Products []Product `gorm:"foreignkey:CategoryID"` // 指定外键为Product中的CategoryID
//}

// Profile 是数据库中的用户个人资料模型
//type Profile struct {
//	gorm.Model
//	UserID   uint
//	Address  string
//	Birthday string
//}
