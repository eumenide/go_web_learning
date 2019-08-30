package main

import (
	"io"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/eumes", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "你好，eumes......")
	})
	mux.HandleFunc("/bye", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "byebye......")
	})
	// 重定向
	mux.HandleFunc("/google", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "http://www.google.com", http.StatusTemporaryRedirect)
	})

	mux.HandleFunc("/hello", sayHello)
	http.ListenAndServe(":8080", mux)
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello world")
}
