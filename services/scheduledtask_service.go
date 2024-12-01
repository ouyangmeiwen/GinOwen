package services

import (
	"GINOWEN/global"
	"GINOWEN/models"
	"fmt"
	"log"
	"time"
)

type ScheduledTaskService struct {
}

func (ScheduledTaskService) getScheduledTasks() (takss []models.ScheduledTask, err error) {
	var ts []models.ScheduledTask
	err = global.OWEN_DB.Model(&models.ScheduledTask{}).Where("status = ? OR status = ?", "pending", "running").Find(&ts).Error
	if err != nil {
		return takss, err
	}
	return ts, err
}

func (ScheduledTaskService) updateTaskStatus(taskID int, status string) error {
	return global.OWEN_DB.Model(&models.ScheduledTask{}).Where("ID=?", taskID).Updates(map[string]interface{}{
		"status": status,
	}).Error
}
func (ScheduledTaskService) updateTaskInternal(taskID int, nexttime time.Time) error {
	return global.OWEN_DB.Model(&models.ScheduledTask{}).Where("ID=?", taskID).Updates(map[string]interface{}{
		"next_run_time": nexttime,
		"status":        "pending",
	}).Error
}
func (ScheduledTaskService) executeTask(task models.ScheduledTask) {
	fmt.Printf("Executing task: %s\n", task.TaskName)
}

// 添加新任务
func (ScheduledTaskService) AddScheduledTask(task models.ScheduledTask) error {
	// 将任务插入数据库
	err := global.OWEN_DB.Create(&task).Error
	if err != nil {
		return fmt.Errorf("failed to add task: %v", err)
	}
	fmt.Printf("Task '%s' added successfully\n", task.TaskName)
	return nil
}

// ProcessTasks 处理并执行所有到期的任务
func (s *ScheduledTaskService) ProcessTasks() error {
	tasks, err := s.getScheduledTasks()
	if err != nil {
		return fmt.Errorf("error getting tasks: %v", err)
	}

	for _, task := range tasks {
		// 如果任务时间已到，执行任务
		if time.Now().After(task.NextRunTime) {
			// 更新任务状态为运行中
			if err := s.updateTaskStatus(task.ID, "running"); err != nil {
				log.Printf("Error updating task %d status to 'running': %v", task.ID, err)
				continue
			}
			// 执行任务
			s.executeTask(task)

			// 更新任务状态为完成
			if err := s.updateTaskStatus(task.ID, "completed"); err != nil {
				log.Printf("Error updating task %d status to 'completed': %v", task.ID, err)
			}

			// 如果任务是周期性的，重新设置下次执行时间
			if task.IntervalSeconds != nil {
				nextRunTime := time.Now().Add(time.Duration(*task.IntervalSeconds) * time.Second)
				if err := s.updateTaskInternal(task.ID, nextRunTime); err != nil {
					log.Printf("Error updating task %d next run time: %v", task.ID, err)
				}
			}
		}
	}

	return nil
}
