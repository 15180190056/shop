package user

import (
	"github.com/gin-gonic/gin"
	"mystopproject/conf"
	"mystopproject/dbserver"
	"net/http"
)

func UpdateUser(c *gin.Context) {
	// 创建一个LoginData实例
	var loginData *conf.User
	// 绑定JSON数据到结构体
	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := dbserver.UpdateUser(loginData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
