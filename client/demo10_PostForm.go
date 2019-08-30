package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

/**
 * @author: eumes
 * @date: 2019-08-30 12:18
 */

func main() {
	urlStr := "http://localhost:8809/login"
	param := url.Values{
		"username": {"eumes"},
		"password": {"123456"},
	}
	response, err := http.PostForm(urlStr, param)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(body))
	}

	fmt.Printf("response: %+v\n", response)
	fmt.Printf("response.Header: %+v\n", response.Header)
	fmt.Printf("response.Cookies: %+v\n", response.Cookies())
}
