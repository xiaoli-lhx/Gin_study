package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 重定向
func main() {
	r := gin.Default()

	r.GET("/index", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
		//c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com")
	})
	r.GET("/a", func(c *gin.Context) {
		//跳转到/b对应的路由处理函数
		c.Request.URL.Path = "/b" //把请求的URL修改
		r.HandleContext(c)        //继续后续的处理
	})
	r.GET("/b", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"hello": "world",
		})
	})
	_ = r.Run(":8080")
}
