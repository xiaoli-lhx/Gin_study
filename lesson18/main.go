package main

//中间件
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// handlerFunc
func indexHandler(c *gin.Context) {
	fmt.Println("index")
	name, ok := c.Get("name") //跨中间件存取值
	if !ok {
		name = "匿名用户"
	}
	c.JSON(http.StatusOK, gin.H{
		"mas": name,
	})
}

// 定义一个中间件m1,统计耗时
func m1(c *gin.Context) {
	fmt.Println("m1 in...")
	//计时
	start := time.Now()
	c.Next() //调用后续的处理函数
	//c.Abort() //阻止调用后续的处理函数
	cost := time.Since(start)
	fmt.Printf("cost:%v\n", cost)
	fmt.Println("m1 out...")
}
func m2(c *gin.Context) {
	fmt.Println("m2 in...")
	c.Set("name", "qimi")
	c.Next() //调用后续的处理函数
	fmt.Println("m2 out...")
}
func authMiddleware(doCheck bool) gin.HandlerFunc {
	//连接数据库
	//或者一些其他准备工作
	return func(c *gin.Context) {
		if doCheck {
			//存放具体的逻辑
			//是否登录的判断
			//if 是登录用户
			//c.Next()
			//else
			//c.Abort()
		} else {
			c.Next()
		}
	}
}
func main() {
	r := gin.Default()
	r.Use(m1, m2, authMiddleware(true)) //全局注册中间件
	r.GET("/index", indexHandler)
	r.GET("/shop", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello world",
		})
	})
	r.GET("/user", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello",
		})
	})
	_ = r.Run(":8080")
}
