package conf

import "gorm.io/gorm"

// Product 对应数据库中的products表
type Product struct {
	gorm.Model
	Code  string
	Price uint
	// 添加CategoryID作为外键
	CategoryID uint `gorm:"column:category_id;ForeignKey:References"`
}
