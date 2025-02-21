package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker %d started processing job %d\n", id, job)
		time.Sleep(time.Second) // 模拟任务处理
		results <- job * 2      // 将结果发送到结果 channel
	}
}

func main() {
	jobs := make(chan int, 5)
	results := make(chan int, 5)

	var wg sync.WaitGroup

	// 启动 3 个 worker goroutine 来并发处理任务
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, jobs, results, &wg)
	}

	// 给 jobs channel 发送任务
	for i := 1; i <= 5; i++ {
		jobs <- i
	}
	close(jobs) // 关闭 jobs channel

	// 等待所有 goroutine 完成
	wg.Wait()

	// 关闭结果 channel 并读取所有结果
	close(results)
	for result := range results {
		fmt.Println("Result:", result)
	}
}
