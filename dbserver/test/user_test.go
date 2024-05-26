package dbserver_test

import (
	"fmt"
	"mystopproject/conf"
	"mystopproject/dbserver"
	"testing"
)

func TestCreateUser(t *testing.T) {
	// 创建一个新用户实例进行测试
	testUser := &conf.User{
		Username: "hetao",
		Password: "03220410Az",
		Email:    "15180190056@example.com",
		Phone:    "456asda",
	}
	// 调用CreateUser函数并检查错误
	dbserver.CreateUser(testUser)

}

func TestGetUserByName(t *testing.T) {
	// 使用db进行测试
	user, err := dbserver.GetUserByName("hetao")
	if err != nil {
		t.Errorf("Failed to get user by ID: %v", err)
	}
	fmt.Println(user.Username)
}

func TestUpdateUser(t *testing.T) {
	// 创建一个新用户实例进行测试
	testUser := &conf.User{
		Username: "hetao",
		Password: "03220410Az",
		Email:    "15180190053@example.com",
		Phone:    "456asda",
	}
	// 调用CreateUser函数并检查错误
	dbserver.UpdateUser(testUser)
}

func TestDeleteByName(t *testing.T) {
	// 尝试删除用户名为"hetao"的用户
	err := dbserver.DeleteByName("hetao")
	if err != nil {
		t.Errorf("删除用户名为'hetao'的用户失败: %v", err)
	}

}
