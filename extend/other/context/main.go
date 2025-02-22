package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	go func() {
		<-time.After(1 * time.Second)
		fmt.Println("Task completed before timeout")
		cancel() // 取消任务
	}()
	select {
	case <-ctx.Done():
		fmt.Println("Timeout reached, cancelling task")
	}
}
