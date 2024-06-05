package conf

import "gorm.io/gorm"

// Product 对应数据库中的products表
type Product struct {
	gorm.Model
	Orders      []Order `gorm:"foreignkey:ProductID"`
	ProductName string
	Price       float64
	Stock       int
}
