package utils

import (
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
