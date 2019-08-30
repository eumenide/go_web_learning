package main

import (
	"eumes.cn/go_web_learning/common"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

/**
 * @author: eumes
 * @date: 2019-08-30 12:01
 */

func main() {
	resp, err := http.Post(common.URL, "application/x-www-form-urlencoded", strings.NewReader("theCityName=南京"))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))
}
