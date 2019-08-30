package main

import (
	"fmt"
	"time"
)

const DATE_FORMAT = "2006-01-02"

func main() {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	t := time.Now().Format(DATE_FORMAT)
	zero, _ := time.ParseInLocation(DATE_FORMAT, t, loc)

	fmt.Println(zero)
}
