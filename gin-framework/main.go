package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello World")

	// 1. 创建路由
	r := gin.Default()

	// 2. 绑定路由规则
	// http://localhost:8080
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})

	// 3. 设置监听端口
	r.Run(":8080")
}
