package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

func main() {
	r := gin.Default()
	//加载静态文件
	r.Static("/xxx", "./statics")
	//gin框架中给模板添加自定义函数
	r.SetFuncMap(template.FuncMap{
		"safe": func(s string) template.HTML {
			return template.HTML(s)
		},
	})
	r.LoadHTMLGlob("templates/**/*") //模板解析
	r.GET("/posts/index", func(c *gin.Context) {
		//HTTP请求
		c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{ //渲染模板
			"title": "liwenzhou.com", //	H is a shortcut for map[string]any
		})
	})
	r.GET("/users/index", func(c *gin.Context) {
		//HTTP请求
		c.HTML(http.StatusOK, "users/index.tmpl", gin.H{ //渲染模板
			"title": "<a href='https://liwenzhou.com'>李文周的博客</a>", //	H is a shortcut for map[string]any
		})
	})
	r.GET("posts/home", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.html", nil)
	})
	_ = r.Run(":8080") //启动server
}
