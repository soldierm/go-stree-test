package main

import (
	"go-stress-test/components"
	"log"
	"net/http"
	"sync"
	"time"
)

var wg sync.WaitGroup
var total int

var url = components.GlobalConfig.Address.UriToString()

func main() {
	times := components.GlobalConfig.Test.Times / components.GlobalConfig.Test.Threads

	now := time.Now().UnixNano()
	log.Printf("%d个线程请求开始..", components.GlobalConfig.Test.Threads)

	for i := 0; i < components.GlobalConfig.Test.Threads; i++ {
		wg.Add(1)
		go groupRequest(times)
	}

	wg.Wait()
	end := time.Now().UnixNano()
	log.Printf("总共请求%d次", total)
	log.Printf("总共耗时%dns", end-now)
}

//循环发送请求
func groupRequest(times int) {
	defer func() {
		total += times
		wg.Done()
	}()
	for i := 0; i < times; i++ {
		sendRequest()
	}
}

//发送请求
func sendRequest() {
	response, err := http.Get(url)
	if err != nil {
		log.Println("something went wrong")
	}
	//这一步必须要做，不然golang会缓存请求
	//@see https://stackoverflow.com/questions/33238518/what-could-happen-if-i-dont-close-response-body-in-golang
	defer response.Body.Close()
}
