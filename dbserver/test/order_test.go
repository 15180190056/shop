package dbserver_test

import (
	"mystopproject/conf"
	"mystopproject/dbserver"
	"testing"
)

func TestCreateOrder(t *testing.T) {
	testOrder := &conf.Order{
		OrderID:     1,
		Quantity:    1,
		UserID:      1,
		ProductID:   1,
		OrderStatus: "已支付",
		TotalPrice:  876,
	}
	// 调用CreateUser函数并检查错误
	dbserver.CreateOrder(testOrder)
}
