package dbserver_test

import (
	"fmt"
	"mystopproject/conf"
	"mystopproject/dbserver"
	"testing"
)

func TestCreateProduct(t *testing.T) {
	// 创建一个新用户实例进行测试
	testProduct := &conf.Product{
		ProductName: "驼奶粉",
		Price:       876,
		Stock:       32,
	}
	// 调用CreateUser函数并检查错误
	dbserver.CreateProduct(testProduct)
}

func TestGetProductByName(t *testing.T) {
	// 使用db进行测试
	product, err := dbserver.GetProductByName("驼奶粉")
	if err != nil {
		t.Errorf("Failed to get user by ID: %v", err)
	}
	fmt.Println(product)
}

func TestUpdateProduct(t *testing.T) {
	// 创建一个新用户实例进行测试
	testProduct := &conf.Product{
		ProductName: "驼奶粉",
		Price:       876,
		Stock:       32,
	}
	// 调用CreateUser函数并检查错误
	dbserver.UpdateProduct(testProduct)
}

func TestDeleteProductByName(t *testing.T) {
	// 尝试删除用户名为"hetao"的用户
	err := dbserver.DeleteProductByName("驼奶粉")
	if err != nil {
		t.Errorf("删除用户名为'hetao'的用户失败: %v", err)
	}

}
