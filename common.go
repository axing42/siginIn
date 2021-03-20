package main

import (
	"bufio"
	"net/http"
	"os"
	"syscall"
)

func Must(e error) {
	if e != nil {
		panic(e)
	}
}
func DClose(body *http.Response) {
	err := body.Body.Close()
	Must(err)
}

func MyCurl(url string) *http.Response {
	get, err := http.Get(url)
	Must(err)
	if get.StatusCode != 200 {
		println("请求失败,状态码: ", get.StatusCode)
		syscall.Exit(0)
	}
	return get
}

// 临时文件操作
func fileOperation(path string) string {
	file, _ := os.Open(path)
	defer func() {
		_ = file.Close()
	}()
	scanner := bufio.NewScanner(file)
	var line string
	for scanner.Scan() {
		line += scanner.Text()
	}
	return line
}
