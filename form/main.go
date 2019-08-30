package main

import (
	"log"
	"net/http"
)

/**
 * @author: eumes
 * @date: 2019-08-30 15:42
 */

func main() {
	http.HandleFunc("/hello", SayHello)
	http.HandleFunc("/login", Login)
	http.HandleFunc("/register", Register)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
