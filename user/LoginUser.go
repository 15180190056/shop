package user

import (
	"github.com/gin-gonic/gin"
	"mystopproject/conf"
	"mystopproject/dbserver"
	"net/http"
)

func LoginUser(c *gin.Context) {
	// 创建一个LoginData实例
	var loginData conf.User
	// 绑定JSON数据到结构体
	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userFrom, err := dbserver.GetUserByName(string(loginData.Username))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	if loginData.Password == userFrom.Password {
		c.JSON(http.StatusOK, gin.H{"login succees": loginData.Username})
	} else {
		c.JSON(http.StatusAccepted, gin.H{"密码错误": loginData.Username})
	}

}
