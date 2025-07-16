package main

//querystring
import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//GET请求URL ?后面是query string参数
	//key=value格式，多个key-value用&连接
	//eg: /web?query=小李&age=20
	r.GET("/web", func(c *gin.Context) {
		//获取浏览器那边发请求携带的query string参数
		//方法一
		name := c.Query("query")
		age := c.Query("age")
		//方法二
		//name:=c.DefaultQuery("query","somebody")	//取不到就用默认值
		//方法三
		//name,ok:=c.GetQuery("query")
		//if !ok{
		//	//取不到
		//	name="somebody"
		//}
		c.JSON(200, gin.H{
			"name": name,
			"age":  age,
		})
	})
	_ = r.Run(":8080")
}
