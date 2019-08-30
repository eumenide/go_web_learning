package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

/**
 * @author: eumes
 * @date: 2019-08-30 12:05
 */

func main() {
	/**
	  请求自己搭建服务器
	*/
	urlStr := "http://localhost:8809/login"
	param := url.Values{
		"username": {"eumes"},
		"password": {"123456"},
	}
	reqBody := bytes.NewBufferString(param.Encode())
	response, err := http.Post(urlStr, "application/x-www-form-urlencoded", reqBody)
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
