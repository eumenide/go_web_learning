package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

/**
他需要2个参数，一个是http.ResponseWriter，另一个是http.Request
往http.ResponseWriter写入什么内容，浏览器的网页源码就是什么内容。
http.Request里面是封装了浏览器发过来的请求（包含路径、浏览器类型等等）。
*/

func sayHello(w http.ResponseWriter, r *http.Request) {
	// w.write([]byte("hello 世界!"))
	fmt.Println("--------------------------------------")

	// 解析参数，默认是不会解析的
	r.ParseForm()

	// 这些信息是输出到服务器端的打印信息
	fmt.Println(r.Form)
	fmt.Println("path: ", r.URL.Path)
	fmt.Println("schema: ", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key: ", k)
		fmt.Println("val: ", strings.Join(v, ""))
	}

	// 这个写入到w的是输出到客户端的
	fmt.Fprintf(w, "Hello GoWeb!")
}

func main ()  {
	/**
		第一个参数： pattern string,
		第二个参数：handler func(ResponseWriter, *Request)
	 */
	// 设置访问的路由
	http.HandleFunc("/", sayHello)

	/**
		第一个参数：addr	监听地址
		第二个参数：handler	通常为空，意味着服务端调用http.DefaultServerMux进行处理，
				而服务端编写的业务逻辑处理程序http.Handle()或http.HandleFunc默认注入到
				http.DefaultServerMux中
	 */
	// 设置监听的端口
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}