package user

import (
	"github.com/gin-gonic/gin"
	"mystopproject/conf"
	"mystopproject/dbserver"
	"net/http"
)

func RegisterUser(c *gin.Context) {
	var userFrom *conf.User
	// 调用CreateUser函数并检查错误
	if err := c.ShouldBindJSON(&userFrom); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := dbserver.CreateUser(userFrom)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
}
