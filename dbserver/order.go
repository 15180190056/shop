package dbserver

import (
	"errors"
	"mystopproject/conf"
	"strings"
)

func CreateOrder(usertable *conf.Order) error {
	DB, err := InitDB()
	if err != nil {
		return err
	}
	DB.Create(&usertable)
	return err
}

func GetUserOrderByName(username string) (*conf.Order, error) {
	var user conf.Order
	DB, err := InitDB()
	if err != nil {
		return nil, err
	}
	DB.Find(&user, "username=?", username)
	return &user, nil // 用户存在，返回用户指针和nil错误
}

func UpdateOrder(order *conf.Order) error {
	DB, err := InitDB()
	if err != nil {
		return err
	}
	DB.Table("users").Where("username = ?", order.OrderID).Updates(order)
	//// UPDATE users SET name='hello', age=18 WHERE id IN (10, 11);
	return nil
}

func DeleteOrderByName(username string) error {
	DB, err := InitDB()
	if err != nil {
		return err
	}
	result := DB.Where("username = ?", username).Delete(&conf.User{})
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "record not found") {
			return errors.New("用户不存在") // 用户不存在，返回自定义错误
		}
		return result.Error
	}
	return nil
}
