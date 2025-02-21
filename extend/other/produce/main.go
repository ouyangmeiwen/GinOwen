package main

import (
	"fmt"
	"time"
)

func produce(ch chan<- int) {
	for i := 0; i < 100; i++ {
		ch <- i
		fmt.Printf("%d加入chan\n", i)
		time.Sleep(time.Second)
	}
	close(ch)
}

func consume(ch <-chan int) {
	for item := range ch {
		fmt.Printf("取出%d\n", item)
		time.Sleep(time.Second * 2)
	}
}

func main() {
	ch := make(chan int, 5)
	go produce(ch)
	go consume(ch)
	time.Sleep(1000 * time.Second)
}
