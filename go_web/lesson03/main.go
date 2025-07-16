package main

import (
	"github.com/gin-gonic/gin"
)

func sayHello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello gin",
	})
}
func main() {
	r := gin.Default() // 返回默认的路由引擎

	//指定用户使用GET请求访问/hello时，执行sayHello这个函数
	r.GET("/hello", sayHello)

	//启动服务
	_ = r.Run(":8080")
}

//func sayHello(w http.ResponseWriter, r *http.Request) {
//	b, _ := os.ReadFile("./hello.txt")
//	_, _ = fmt.Fprintln(w, string(b))
//}
//
//func main() {
//	http.HandleFunc("/hello", sayHello)
//	err := http.ListenAndServe(":8080", nil)
//	if err != nil {
//		fmt.Printf("http serve failed,err:%v\n", err)
//		return
//	}
//}
