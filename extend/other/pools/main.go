package main

import (
	"fmt"
	"sync"
)

type Task struct {
	ID   int
	Name string
}

// sync.Pool 是线程安全的，因为它确保每个 goroutine 在使用池中对象时是独占的，其他 goroutine 在该对象被 Put() 之前不能访问或修改它。
// 即使多个 goroutine 在同一时间访问池中的对象，每个 goroutine 获取到的对象是独立的，并且在对象被放回池之前，其他 goroutine 不能修改该对象，从而保证了数据的一致性和安全性。
func main() {
	var taskPool sync.Pool
	taskPool.New = func() interface{} {
		return &Task{} // 每次从池中获取不到对象时，创建新的对象
	}

	var wg sync.WaitGroup
	numWorkers := 5
	numTasks := 10

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()

			// 每个worker处理不同的任务
			for j := workerID; j < numTasks; j += numWorkers {
				task := taskPool.Get().(*Task) // 从池中获取任务对象

				task.ID = j
				task.Name = fmt.Sprintf("Task %d", j)

				fmt.Printf("Worker %d is processing %s\n", workerID, task.Name)

				taskPool.Put(task) // 放回池中
			}
		}(i)
	}

	wg.Wait()
	fmt.Println("All tasks processed")
}
