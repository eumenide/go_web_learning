package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

/**
 * @author: eumes
 * @date: 2019-08-30 14:31
 */

func main() {
	http.HandleFunc("/hello", sayHello)
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	// 解析URL传递的参数，对于POST则解析响应包的主体
	r.ParseForm()
	// 注意：如果没有调用ParseForm方法，下面无法获取表单的数据
	fmt.Println(r.Form)
	fmt.Println("path: ", r.URL.Path)
	fmt.Println("schema: ", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])

	for k, v := range r.Form {
		fmt.Println("key: ", k)
		fmt.Println("value: ", strings.Join(v, ""))
	}

	// 这个写入到w的是输出到客户端的
	fmt.Fprintf(w, "Hello my route!")
}

func login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println("method: ", r.Method)
	if r.Method == http.MethodGet {
		t, _ := template.ParseFiles("html/login.html")
		t.Execute(w, nil)
	} else {
		// 请求的是登录数据，那么执行登录的逻辑判断
		fmt.Println("username: ", r.Form["username"])
		fmt.Println("password: ", r.Form["password"])
	}
}
