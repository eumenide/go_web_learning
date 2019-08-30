package main

import (
	"bytes"
	"eumes.cn/go_web_learning/common"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

/**
 * @author: eumes
 * @date: 2019-08-30 11:55
 */

func main() {
	client := http.Client{}
	param := &url.Values{
		"theCityName": {"南京"},
	}
	reqBody := bytes.NewBufferString(param.Encode())

	response, err := client.Post(common.URL, "application/x-www-form-urlencoded", reqBody)
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

	fmt.Printf("%+v\n", response)
}
