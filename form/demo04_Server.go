package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

/**
 * @author: eumes
 * @date: 2019-08-30 16:11
 */
func main() {
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method: ", r.Method)
	if r.Method == http.MethodGet {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))
		fmt.Println("token--->", token)
		t, _ := template.ParseFiles("html/test.gtpl")
		t.Execute(w, token)
	} else {
		r.ParseForm()
		token := r.Form.Get("token")
		if token != "" {
			fmt.Println("token: ", token)
		} else {
			fmt.Println("token有误。。")
		}

		fmt.Println("username length: ", len(r.Form["username"][0]))
		fmt.Println("username: ", template.HTMLEscapeString(r.Form.Get("username"))) //输出到服务器端
		fmt.Println("password: ", template.HTMLEscapeString(r.Form.Get("password")))
		template.HTMLEscape(w, []byte(r.Form.Get("username"))) //输出到客户端
	}
}
