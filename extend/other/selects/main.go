package main

import (
	"fmt"
	"time"
)

func task(ch chan<- string) {
	time.Sleep(2 * time.Second)
	ch <- "Task completed"
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	// 启动两个任务
	go task(ch1)
	go task(ch2)

	// 使用 select 等待多个 channel
	select {
	case msg := <-ch1:
		fmt.Println("Received from ch1:", msg)
	case msg := <-ch2:
		fmt.Println("Received from ch2:", msg)
	case <-time.After(3 * time.Second):
		fmt.Println("Timeout occurred")
	}
}
