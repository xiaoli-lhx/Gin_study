package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/json", func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{
			"name":    "小王子",
			"massage": "hello world",
			"age":     18,
		})
	})
	//方法二结构体，灵活使用tag来对结构体字段进行定制化操作
	type msg struct {
		Name    string `json:"name"`
		Massage string
		Age     int
	}
	r.GET("/other_json", func(c *gin.Context) {
		data := msg{
			"小王子",
			"hello golang",
			18,
		}
		c.JSON(http.StatusOK, data)
	})
	_ = r.Run(":8080")
}
