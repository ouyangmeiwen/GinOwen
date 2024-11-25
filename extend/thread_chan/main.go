package main

import (
	"fmt"
	"strconv"
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

func Work(taskid string, ch_send <-chan Task, ch_work chan<- TaskResult, done chan<- bool) {
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
	// 当当前 worker 退出时，通知任务完成
	done <- true
}

func main() {
	task_count := 100 // 增加通道缓冲区大小
	thread_count := 4

	ch_send := make(chan Task, task_count)
	ch_work := make(chan TaskResult, task_count)
	done := make(chan bool) // 用于追踪 worker 完成情况

	// 启动多个 worker goroutine 来处理任务
	for i := 0; i < thread_count; i++ {
		go Work(strconv.Itoa(i), ch_send, ch_work, done)
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

	// 另一个 goroutine 等待所有 worker 完成
	go func() {
		for i := 0; i < thread_count; i++ {
			<-done // 等待每个 worker 发送完成信号
		}
		close(ch_work) // 当所有 worker 完成后，关闭结果通道
	}()

	// 读取并打印任务结果
	for result := range ch_work {
		fmt.Printf("ID %s, Success %t, Msg %s\n", result.Id, result.Success, result.Msg)
	}

	fmt.Println("打印结束")
}
