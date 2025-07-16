package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func f1(w http.ResponseWriter, r *http.Request) {
	//自定义函数
	//要么只有一个返回值，要么第二个返回值是error
	k := func(name string) (string, error) {
		return name + "年轻又帅气", nil
	}
	//定义模板
	//解析模板
	//告诉模板引擎，多了一个自定义函数kua，必须在解析模板之前
	t, err := template.New("f.tmpl").Funcs(template.FuncMap{
		"kua": k,
	}).ParseFiles("./f.tmpl")

	if err != nil {
		fmt.Printf("Parse template failed,err:%v", err)
		return
	}
	name := "小王子"
	//渲染模板
	err = t.Execute(w, name)
	if err != nil {
		fmt.Printf("render template failed,err:%v", err)
		return
	}

}

func demo1(w http.ResponseWriter, r *http.Request) {
	//解析模板
	t, err := template.ParseFiles("t.tmpl", "ul.tmpl")
	if err != nil {
		fmt.Printf("Parse template failed,err:%v", err)
		return
	}
	name := "小王子"
	err = t.Execute(w, name)
	if err != nil {
		fmt.Printf("render template failed,err:%v", err)
		return
	}
}

func main() {
	http.HandleFunc("/", f1)
	http.HandleFunc("/tmplDemo", demo1)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("HTTP server start failed,err: %v\n", err)
		return
	}
}
