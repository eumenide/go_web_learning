package main

import (
	"fmt"
	"net/http"
)

type Mymux struct {
}

// 自定义简易路由
func (p *Mymux) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	if r.URL.Path == "/" {
		sayHello(w, r)
		return
	}

	http.NotFound(w, r)

	return
}

func sayHello(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Hello my route!")
}

func main() {
	mux := &Mymux{}
	http.ListenAndServe(":8080", mux)
}