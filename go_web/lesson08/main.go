package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("index.tmpl").
		Delims("{[", "]}").
		ParseFiles("index.tmpl")
	if err != nil {
		fmt.Printf("parse template failed,err:%v", err)
		return
	}
	name := "小王子"
	_ = t.Execute(w, name)
}
func xss(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("xss.tmpl").Funcs(template.FuncMap{
		"safe": func(s string) template.HTML {
			return template.HTML(s)
		},
	}).ParseFiles("xss.tmpl")
	if err != nil {
		fmt.Printf("parse template failed,err:%v", err)
	}
	str1 := "<script>alert('hello')</script>"
	str2 := "<a href='https://www.liwenzhou.com'>liwenzhou的博客</a>"
	_ = t.Execute(w, map[string]string{
		"str1": str1,
		"str2": str2,
	})
}
func main() {
	http.HandleFunc("/index", index)
	http.HandleFunc("/xss", xss)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("http listen failed,err:%v", err)
		return
	}
}
