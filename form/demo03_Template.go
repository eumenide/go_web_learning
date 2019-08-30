package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

/**
 * @author: eumes
 * @date: 2019-08-30 15:54
 */

func main() {
	http.HandleFunc("/login", login)
	http.HandleFunc("/login2", login2)
	http.HandleFunc("/login3", login3)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	if r.Method == http.MethodGet {
		t, _ := template.ParseFiles("html/login.html")
		t.Execute(w, nil)
	} else {
		// 请求的是登录数据，那么执行登录的逻辑判断
		username := r.Form.Get("username")

		fmt.Println("username: 	", template.HTMLEscapeString(username))
		fmt.Println("password: ", template.HTMLEscapeString(r.Form.Get("password")))

		template.HTMLEscape(w, []byte(username))
	}
}

func login2(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if r.Method == http.MethodGet {
		t, _ := template.ParseFiles("html/login2.html")
		t.Execute(w, nil)
	} else {
		// 请求的是登录数据，那么执行登录的逻辑判断
		username := r.Form.Get("username")
		fmt.Println(username)

		// 进行模板解析
		t, err := template.New("foo").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
		err = t.ExecuteTemplate(w, "T", username)

		// 如果转义失败 抛出对应错误 终止程序
		if err != nil {
			log.Fatal(err)
		}
	}
}

func login3(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if r.Method == http.MethodGet {
		t, _ := template.ParseFiles("html/login3.html")
		t.Execute(w, nil)
	} else {
		// 请求的是登录数据，那么执行登录的逻辑判断
		username := r.Form.Get("username")
		fmt.Println(username)

		// 进行模板解析
		t, err := template.New("foo").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
		err = t.ExecuteTemplate(w, "T", template.HTML(username))

		// 如果转义失败 抛出对应错误 终止程序
		if err != nil {
			log.Fatal(err)
		}
	}
}
