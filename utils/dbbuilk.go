package utils

import (
	"fmt"
	"log"
	"sync"

	"gorm.io/gorm"
)

// 切片传递的是引用
func batchInsertDone[T any](db *gorm.DB, data []T, batchSize int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < len(data); i += batchSize {
		end := i + batchSize
		if end > len(data) {
			end = len(data)
		}
		batch := data[i:end]
		if err := db.Create(&batch).Error; err != nil {
			log.Printf("Error inserting batch: %v\n", err)
		}
	}
}
func BatchInsert[T any](db *gorm.DB, wg *sync.WaitGroup, datas []T, batchSize int) {
	for i := 0; i < len(datas); i += batchSize {
		end := i + batchSize
		if end > len(datas) {
			end = len(datas)
		}
		wg.Add(1)
		go batchInsertDone(db, (datas)[i:end], batchSize, wg)
	}
}

// 处理每个批次的插入操作
func batchChanInsertDone[T any](db *gorm.DB, data []T, wg *sync.WaitGroup, errCh chan<- error) {
	defer wg.Done()
	if err := db.Create(&data).Error; err != nil {
		errCh <- err // 将错误通过通道发送给主程序
		log.Printf("Error inserting batch: %v\n", err)
	}
}

// 批量插入函数，支持限制最大并发数
func BatchChanInsert[T any](db *gorm.DB, datas []T, batchSize int, maxGoroutines int) error {
	if batchSize <= 0 {
		return fmt.Errorf("invalid batchSize: must be greater than 0")
	}

	if len(datas) == 0 {
		return nil
	}

	var wg sync.WaitGroup
	errCh := make(chan error, len(datas)/batchSize+1)

	// 使用信号量控制并发数
	sem := make(chan struct{}, maxGoroutines)

	// 遍历数据，分批次插入
	for i := 0; i < len(datas); i += batchSize {
		end := i + batchSize
		if end > len(datas) {
			end = len(datas)
		}
		batch := datas[i:end]

		// 添加到等待组
		wg.Add(1)

		// 获取信号量，限制最大并发数
		sem <- struct{}{}

		// 启动 goroutine 执行批次插入
		go func(batch []T) {
			defer func() { <-sem }() // 完成后释放信号量
			batchChanInsertDone(db, batch, &wg, errCh)
		}(batch)
	}

	// 等待所有 goroutine 完成
	wg.Wait()

	// 关闭错误通道
	close(errCh)

	// 如果有错误，返回第一个错误
	if err := <-errCh; err != nil {
		return err
	}

	return nil
}
