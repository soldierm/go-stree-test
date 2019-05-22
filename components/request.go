package components

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
)

var wg sync.WaitGroup
var total int
var url string
var handle func() (resp *http.Response, err error)

//初始化操作
func InitRequest() {
	InitConfig()
	switch GlobalConfig.Method.Type {
	case "get":
		handle = get
	case "post":
		handle = post
	}
}

//开始跑喽
func Start() {
	url = GlobalConfig.Address.UriToString()
	times := GlobalConfig.Test.Times / GlobalConfig.Test.Threads
	now := time.Now().UnixNano()
	output(fmt.Sprintf("请求地址：%s..", url))
	output(fmt.Sprintf("%d个线程请求开始..", GlobalConfig.Test.Threads))

	for i := 0; i < GlobalConfig.Test.Threads; i++ {
		wg.Add(1)
		go groupRequest(times)
	}

	wg.Wait()
	end := time.Now().UnixNano()
	output(fmt.Sprintf("总共请求%d次..", total))
	output(fmt.Sprintf("总共耗时%dns..", end-now))
}

//循环发送请求
func groupRequest(times int) {
	for i := 0; i < times; i++ {
		sendRequest()
	}
	total += times
	defer wg.Done()
}

//发送请求
func sendRequest() {
	response, err := handle()
	if err != nil {
		output("something went wrong")
	}
	//这一步必须要做，不然golang会缓存请求
	//@see https://stackoverflow.com/questions/33238518/what-could-happen-if-i-dont-close-response-body-in-golang
	defer response.Body.Close()
}

//普通get请求
func get() (resp *http.Response, err error) {
	return http.Get(url)
}

//普通post请求，需要加上body内容
func post() (resp *http.Response, err error) {
	contentType := GlobalConfig.Method.ContentType
	var body io.Reader
	/* 判断ContentType，分别读取不同的结构作为body */
	if GlobalConfig.Method.ContentType == "application/json" {
		jsonByte, _ := json.Marshal(GlobalConfig.Method.JsonBody)
		body = bytes.NewBuffer(jsonByte)
	} else {
		body = strings.NewReader(GlobalConfig.Method.FormBody.Encode())
	}
	return http.Post(url, contentType, body)
}

func output(data string) {
	if isTerminal() {
		log.Println(data)
	}
}
