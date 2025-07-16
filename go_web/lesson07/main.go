package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("index.tmpl")
	if err != nil {
		fmt.Printf("parse template failed,err:%v", err)
		return
	}
	msg := "小王子"
	err = t.Execute(w, msg)
	if err != nil {
		fmt.Printf("render template failed,err:%v", err)
		return
	}
}
func home(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("home.tmpl")
	if err != nil {
		fmt.Printf("parse template failed,err:%v", err)
		return
	}
	msg := "小王子"
	err = t.Execute(w, msg)
	if err != nil {
		fmt.Printf("render template failed,err:%v", err)
		return
	}
}
func index2(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./template/base.tmpl", "./template/index2.tmpl")
	if err != nil {
		fmt.Printf("parse template failed,err:%v", err)
	}
	name := "小王子"
	err = t.ExecuteTemplate(w, "index2.tmpl", name)
	if err != nil {
		fmt.Printf("render template failed,err:%v", err)
		return
	}
}
func home2(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./template/base.tmpl", "./template/home2.tmpl")
	if err != nil {
		fmt.Printf("parse template failed,err:%v", err)
	}
	name := "七米"
	err = t.ExecuteTemplate(w, "home2.tmpl", name)
	if err != nil {
		fmt.Printf("render template failed,err:%v", err)
		return
	}
}
func main() {
	http.HandleFunc("/index", index)
	http.HandleFunc("/home", home)
	http.HandleFunc("/index2", index2)
	http.HandleFunc("/home2", home2)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("http listen failed,err:%v", err)
		return
	}
}
