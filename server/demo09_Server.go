package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

/**
 * @author: eumes
 * @date: 2019-08-30 12:08
 */

func main() {
	server := http.NewServeMux() // 路由
	server.HandleFunc("/login", login)
	http.ListenAndServe(":8809", server)
}

func login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	body, err := ioutil.ReadAll(r.Body)
	r.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("body: ", string(body))
	fmt.Printf("r: %+v\n", r)
	fmt.Printf("%+v\n", r.Header)
	fmt.Printf("%+v\n", r.Header["Content-Type"])
	fmt.Printf("%+v\n", r.Cookies())

	if len(r.Form["username"]) > 0 {
		username := r.Form["username"][0]
		fmt.Println("username: ", username)
	}
	if len(r.Form["password"]) > 0 {
		password := r.Form["password"][0]
		fmt.Println("password: ", password)
	}

	io.WriteString(w, "登录成功")
}
