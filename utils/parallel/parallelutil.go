package parallel

import (
	"GINOWEN/models"
	"sort"
	"sync"
	"time"
)

// 并发 Map
func ParallelMap[T any, R any](items []T, workers int, mapper func(T) R) []R {
	out := make(chan R, len(items))
	var wg sync.WaitGroup

	blockSize := (len(items) + workers - 1) / workers

	for i := 0; i < workers; i++ {
		start := i * blockSize
		end := start + blockSize
		if end > len(items) {
			end = len(items)
		}
		if start >= len(items) {
			break
		}
		wg.Add(1)
		go func(sub []T) {
			defer wg.Done()
			for _, item := range sub {
				out <- mapper(item)
			}
		}(items[start:end])
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	var results []R
	for item := range out {
		results = append(results, item)
	}
	return results
}

// 增加智能判定
func SmartFilter[T any](items []T, workers int, threshold int, filter func(T) bool) []T {
	if len(items) < threshold || workers == 1 {
		// 串行处理
		var result []T
		for _, item := range items {
			if filter(item) {
				result = append(result, item)
			}
		}
		return result
	}
	// 并发处理
	return ParallelFilter(items, workers, filter)
}

// 并发 Filter
func ParallelFilter[T any](items []T, workers int, filter func(T) bool) []T {
	out := make(chan T, len(items))
	var wg sync.WaitGroup

	blockSize := (len(items) + workers - 1) / workers

	for i := 0; i < workers; i++ {
		start := i * blockSize
		end := start + blockSize
		if end > len(items) {
			end = len(items)
		}
		if start >= len(items) {
			break
		}
		wg.Add(1)
		go func(sub []T) {
			defer wg.Done()
			for _, item := range sub {
				if filter(item) {
					out <- item
				}
			}
		}(items[start:end])
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	var results []T
	for item := range out {
		results = append(results, item)
	}
	return results
}

// 并发 Reduce（每个 goroutine 归约一块，然后汇总）
func ParallelReduce[T any, R any](items []T, workers int, zero R, reducer func(R, T) R, combiner func(R, R) R) R {
	var wg sync.WaitGroup
	results := make(chan R, workers)

	blockSize := (len(items) + workers - 1) / workers

	for i := 0; i < workers; i++ {
		start := i * blockSize
		end := start + blockSize
		if end > len(items) {
			end = len(items)
		}
		if start >= len(items) {
			break
		}
		wg.Add(1)
		go func(sub []T) {
			defer wg.Done()
			acc := zero
			for _, item := range sub {
				acc = reducer(acc, item)
			}
			results <- acc
		}(items[start:end])
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	final := zero
	for r := range results {
		final = combiner(final, r)
	}
	return final
}

// 排序（串行，但可以与上面组合使用）
func ParallelSort[T any](items []T, less func(i, j int) bool) {
	sort.Slice(items, less)
}

func Test() {

	var rows []models.Librow // 假设 models.Librow 是你定义的结构体
	// Filter 示例：最近一个月的 RowType=2、RowNo=1
	filtered := ParallelFilter(rows, 8, func(r models.Librow) bool {
		return r.RowNo == 1 && r.RowType == 2 && r.CreationTime.After(time.Now().AddDate(0, -1, 0))
	})

	// Map 示例：提取创建时间戳
	timestamps := ParallelMap(filtered, 8, func(r models.Librow) int64 {
		return r.CreationTime.Unix()
	})
	println(timestamps)
	// Reduce 示例：计算所有创建时间的总和（秒）
	totalTime := ParallelReduce(filtered, 8, int64(0),
		func(acc int64, r models.Librow) int64 {
			return acc + r.CreationTime.Unix()
		},
		func(a, b int64) int64 {
			return a + b
		},
	)
	println(totalTime)
	// Sort 示例：按创建时间降序排列
	ParallelSort(filtered, func(i, j int) bool {
		return filtered[i].CreationTime.After(filtered[j].CreationTime)
	})
}
