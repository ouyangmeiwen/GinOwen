package main

import (
	"fmt"
	"net/http"
	"time"
)

var (
	semaphore     chan struct{}
	maxConcurrent = 500
)

func init() {
	semaphore = make(chan struct{}, maxConcurrent)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	// 获取信号量，限制并发数
	semaphore <- struct{}{}
	defer func() { <-semaphore }()

	// 模拟耗时操作
	time.Sleep(time.Second)

	// 处理请求逻辑
	fmt.Fprintf(w, "Hello, world!\n")
}

func main() {
	http.HandleFunc("/", handleRequest)
	server := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	fmt.Println("Server listening on port 8080")
	if err := server.ListenAndServe(); err != nil {
		fmt.Println("Server failed:", err)
	}
}
