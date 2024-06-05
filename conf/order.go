package conf

import (
	"gorm.io/gorm"
)

// Order 是数据库中的订单模型
type Order struct {
	gorm.Model
	OrderID     uint
	UserID      uint
	ProductID   uint
	Quantity    int
	OrderStatus string
	TotalPrice  float64
}
