package main

import (
	"fmt"
	"sync"
	"time"
)

func produce(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done() // 生产者完成时调用 Done()

	for i := 0; i < 10; i++ {
		ch <- i
		fmt.Printf("生产者 %d 加入 %d\n", i, i)
		time.Sleep(time.Second)
	}
}

func consume(ch <-chan int, wg *sync.WaitGroup, id int) {
	defer wg.Done() // 消费者完成时调用 Done()

	for item := range ch {
		fmt.Printf("消费者 %d 消费了 %d\n", id, item)
		time.Sleep(2 * time.Second)
	}
}

func main() {
	ch := make(chan int, 5)
	var wg sync.WaitGroup

	// 启动多个生产者
	for i := 0; i < 3; i++ { // 3个生产者
		wg.Add(1)
		go produce(ch, &wg)
	}

	// 启动多个消费者
	for i := 0; i < 2; i++ { // 2个消费者
		wg.Add(1)
		go consume(ch, &wg, i+1)
	}

	// 等待所有生产者完成
	wg.Wait()

	// 关闭 channel（在所有生产者完成后关闭）
	close(ch)

	// 给消费者时间处理所有数据
	time.Sleep(5 * time.Second)
}
