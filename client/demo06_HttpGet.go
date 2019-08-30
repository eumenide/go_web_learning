package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

/**
 * @author: eumes
 * @date: 2019-08-30 11:48
 */

func main() {
	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		//handle error
		log.Fatal(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		log.Fatal(err)
	}

	fmt.Println(string(body))
}
