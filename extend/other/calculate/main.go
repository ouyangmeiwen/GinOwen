package main

import (
	"fmt"
	"math"
	"sync"
)

func calculateSquareRoot(x float64, result chan<- float64, wg *sync.WaitGroup) {
	defer wg.Done()
	result <- math.Sqrt(x)
}

func main() {
	var wg sync.WaitGroup
	results := make(chan float64, 5)

	// 启动多个计算任务
	numbers := []float64{4, 16, 25, 36, 49}
	for _, num := range numbers {
		wg.Add(1)
		go calculateSquareRoot(num, results, &wg)
	}

	// 等待所有任务完成
	wg.Wait()
	close(results)

	// 读取结果并输出
	for result := range results {
		fmt.Println("Square root:", result)
	}
}
