package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// http://localhost:8080/user/zhangsan/play
	r.GET("/user/:name/*action", func(c *gin.Context) {
		// 业务逻辑
		name := c.Param("name")     // zhangsan
		action := c.Param("action") // /play
		action = strings.Trim(action, "/")
		c.String(http.StatusOK, name+"\t"+action)
	})

	r.Run(":8080")
}
