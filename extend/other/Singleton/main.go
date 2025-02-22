package main

import (
	"fmt"
	"sync"
)

// 单例结构体
type Singleton struct {
	Value int
}

var instance *Singleton
var once sync.Once

// 获取单例实例
func GetInstance() *Singleton {
	once.Do(func() {
		// 在第一次调用时初始化实例
		instance = &Singleton{Value: 42}
	})
	return instance
}

func main() {
	// 使用 goroutines 并发调用 GetInstance，确保只创建一个实例
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			singleton := GetInstance()
			fmt.Printf("Goroutine %d: Singleton Value = %d\n", i, singleton.Value)
			// 打印单例的地址和它的值
			fmt.Printf("Goroutine %d: Singleton Address = %p, Singleton Value = %d\n", i, singleton, singleton.Value)
		}(i)
	}

	// 等待所有 goroutines 执行完毕
	wg.Wait()
}
