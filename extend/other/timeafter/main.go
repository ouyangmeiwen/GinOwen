package main

import (
	"fmt"
	"time"
)

func longRunningTask(ch chan<- string) {
	time.Sleep(5 * time.Second)
	ch <- "Task finished"
}

func main() {
	ch := make(chan string)

	go longRunningTask(ch)

	// 使用 select 处理超时
	select {
	case result := <-ch:
		fmt.Println(result)
	case <-time.After(3 * time.Second):
		fmt.Println("Timeout occurred")
	}
}
