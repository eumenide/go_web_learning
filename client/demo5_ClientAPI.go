package main

import (
	"../common"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	/**
	方式二 调用client的API
		1. 先生成client
		2. 之后用client.get/post..

	还额外添加了req.Header.Set("Content-Type", bodyType)...
	*/

	// step1：创建client对象
	client := http.Client{}

	// step2：调用get方法，发送请求
	response, err := client.Get(common.URL)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	if response.StatusCode == 200 {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(body))
	}

	fmt.Printf("%+v\n", response)

}
