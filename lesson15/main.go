package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
)

// 文件上传
func main() {
	r := gin.Default()
	r.LoadHTMLFiles("index.html")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	r.POST("/update", func(c *gin.Context) {
		//从请求中读取文件
		f, err := c.FormFile("f1")
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{
				"error": err.Error(),
			})
		} else {
			//将读取到的文件保存在本地(服务端本地)
			dst := path.Join("./", f.Filename)
			_ = c.SaveUploadedFile(f, dst)
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
		}
	})
	_ = r.Run(":8080")
}
