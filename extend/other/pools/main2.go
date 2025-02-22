package main

import (
	"fmt"
	"sync"
)

// ，我们使用一个共享的 Task 对象。所有 goroutine 都会共享这个对象，每次处理任务时都修改该对象。这会引入并发问题，因为多个 goroutine 可能会同时修改共享对象的字段。
func main2() {
	var task Task   // 共享的任务对象
	numTasks := 5   // 模拟5个任务
	numWorkers := 3 // 使用3个goroutine并发处理任务

	var wg sync.WaitGroup
	for w := 0; w < numWorkers; w++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for i := workerID; i < numTasks; i += numWorkers {
				// 共享的任务对象，每次都修改它
				task.ID = i
				task.Name = fmt.Sprintf("Task %d", i)
				// 处理任务
				fmt.Printf("Worker %d is processing %s\n", workerID, task.Name)

				// 注意：任务对象没有复用，不存在池管理
			}
		}(w)
	}

	// 等待所有goroutine完成
	wg.Wait()
	fmt.Println("All tasks processed")
}
