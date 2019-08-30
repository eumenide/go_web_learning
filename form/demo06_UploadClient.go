package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

/**
 * @author: eumes
 * @date: 2019-08-30 16:39
 */
func main() {
	targetUrl := "http://localhost:8080/upload"
	fileName := "./README.md"
	postFile(fileName, targetUrl)
}

func postFile(fileName, targetUrl string) error {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	// 关键的一些操作
	fileWriter, err := bodyWriter.CreateFormFile("uploadfile", fileName)
	if err != nil {
		fmt.Println("error writing to buffer")
		return err
	}

	// 打开文件句柄操作
	fh, err := os.Open(fileName)
	if err != nil {
		fmt.Println("error opening file")
		return err
	}

	// iocopy
	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		fmt.Println("error iocopy file")
		return err
	}
	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()
	resp, err := http.Post(targetUrl, contentType, bodyBuf)
	if err != nil {
		fmt.Println("error post")
		return err
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error reading response body")
		return err
	}
	fmt.Println(resp.Status)
	fmt.Println(string(respBody))
	return nil
}
