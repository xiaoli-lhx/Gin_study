package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	//访问/index的GET请求会走这一条处理逻辑
	//路由
	//获取
	r.GET("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "GET",
		})
	})
	//创建
	r.POST("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "POST",
		})
	})
	//更新（局部）
	r.PUT("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "PUT",
		})
	})
	//删除
	r.DELETE("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "DELETE",
		})
	})
	//处理所有
	r.Any("/user", func(c *gin.Context) {
		switch c.Request.Method {
		case "GET":
			c.JSON(http.StatusOK, gin.H{"method": "GET"})
		case http.MethodPost:
			c.JSON(http.StatusOK, gin.H{"method": "POST"})
		}
	})
	//NoRoute访问不存在的
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"msg": "liwenzhou.com"})
	})
	//路由组
	//把公共的前缀提取出来创建一个路由组
	videoGroup := r.Group("/video")
	{
		videoGroup.GET("index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"method": "index"})
		})
		videoGroup.GET("/xx", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"method": "xx"})
		})
		videoGroup.GET("/oo", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"method": "oo"})
		})
		//嵌套
		ooGroup := r.Group("/oo")
		ooGroup.GET("xx", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"method": "xx"})
		})
	}
	_ = r.Run(":8080")
}
