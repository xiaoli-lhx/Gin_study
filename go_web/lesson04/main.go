package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	//2.解析模板
	t, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Printf("Parse template failed,err:%v", err)
		return
	}
	//3.渲染模板
	name := "小王子"
	err = t.Execute(w, name)
	if err != nil {
		fmt.Printf("render template failed,err:%v", err)
	}
}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("HTTP server start failed,err: %v\n", err)
		return
	}

}
