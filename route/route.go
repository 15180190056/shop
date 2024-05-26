package route

import (
	"github.com/gin-gonic/gin"
	"mystopproject/user"
	"net/http"
)

func Route() {
	route := gin.Default()
	// 设置CORS中间件
	route.Use(func(c *gin.Context) {
		// 允许所有的域进行跨域请求
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

		// 允许的方法
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

		// 预检间隔时间
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		// 允许发送带凭证的请求
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		// 处理OPTIONS请求
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		} else {
			c.Next()
		}
	})

	//用户管理
	usermanager := route.Group("user")
	{
		// 注册新用户
		usermanager.POST("/RegisterUser", user.RegisterUser)
		// 用户登录
		usermanager.POST("/LoginUser", user.LoginUser)
		// 更新用户信息
		usermanager.POST("/UpdateUser", user.UpdateUser)
		// 删除用户
		usermanager.POST("/DeleteUser", user.DeleteUser)

	}

	////商品管理
	//goods := route.Group("goods")
	//{
	//	//添加商品信息
	//	goods.POST("/AddProduct", tmp.AddCategory)
	//	//商品编辑
	//	goods.POST("/EditProduct", tmp.EditProduct)
	//	//商品上架
	//	goods.POST("/ShelveProduct", tmp.ShelveProduct)
	//	//商品下架
	//	goods.POST("/UnshelveProduct", tmp.UnshelveProduct)
	//}
	//
	////订单管理
	//order := route.Group("order")
	//{
	//	// 打印所有订单信息
	//	order.GET("list", src.PrintAllOrders)
	//	//修改订单信息
	//	order.GET("list", src.UpdateOrderStatus)
	//	// 查询订单
	//	order.GET("list", src.GetOrder)
	//	// 删除订单  DeleteOrder
	//	order.GET("list", src.DeleteOrder)
	//	// 添加订单  AddOrder
	//	order.GET("list", src.AddOrder)
	//}

	route.Run(":8083")
}
