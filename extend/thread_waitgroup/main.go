package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

type Task struct {
	Params string
	Id     string
}

type TaskResult struct {
	Id      string
	Success bool
	Msg     string
}

func Work(taskid string, ch_send <-chan Task, ch_work chan<- TaskResult, wg *sync.WaitGroup) {
	defer wg.Done() // 确保 goroutine 完成任务后调用 Done
	for request := range ch_send {
		fmt.Printf("Worker %s 开始接收到任务ID:%s 参数:%s\n", taskid, request.Id, request.Params)
		// 模拟任务处理
		time.Sleep(1 * time.Second)
		// 生成任务结果
		result := TaskResult{
			Success: true,
			Msg:     request.Params,
			Id:      request.Id,
		}
		// 发送结果
		ch_work <- result
		fmt.Printf("Worker %s 已处理完任务ID:%s\n", taskid, request.Id)
	}
}

func main() {
	task_count := 100
	thread_count := 4
	ch_send := make(chan Task, task_count)
	ch_work := make(chan TaskResult, task_count)
	var wg sync.WaitGroup

	// 启动更多 worker goroutine 来处理任务
	for i := 0; i < thread_count; i++ { // 增加 worker 数量到 4
		wg.Add(1)
		go Work(strconv.Itoa(i), ch_send, ch_work, &wg)
	}

	// 发送任务
	go func() {
		for i := 0; i < task_count; i++ {
			var t Task
			t.Id = strconv.Itoa(i)
			t.Params = "参数" + t.Id
			ch_send <- t
		}
		close(ch_send) // 关闭任务通道，表示不再发送任务
	}()

	// 开启一个 goroutine 等待所有任务完成后关闭 ch_work 通道
	go func() {
		wg.Wait()
		close(ch_work) // 确保所有 worker 完成任务后关闭 ch_work
	}()

	// 读取并打印任务结果
	for result := range ch_work {
		fmt.Printf("ID %s, Success %t, Msg %s\n", result.Id, result.Success, result.Msg)
	}

	fmt.Println("打印结束")
}
