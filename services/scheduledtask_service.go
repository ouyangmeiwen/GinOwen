package services

import (
	"GINOWEN/global"
	"GINOWEN/models"
	"GINOWEN/models/request"
	"GINOWEN/models/response"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

type ScheduledTaskService struct {
}

func (ScheduledTaskService) getScheduledTasks() (takss []models.ScheduledTask, err error) {
	if global.OWEN_CONFIG.System.TaskDB == "redis" {
		var tasks []models.ScheduledTask
		keys, err := global.OWEN_REDIS.Keys(global.Ctx, "task:*").Result() // 获取所有任务的键
		if err != nil {
			return nil, fmt.Errorf("error getting task keys from Redis: %v", err)
		}

		for _, key := range keys {
			taskData, err := global.OWEN_REDIS.HGetAll(global.Ctx, key).Result()
			if err != nil {
				log.Printf("Error fetching task data for %s: %v", key, err)
				continue
			}

			var task models.ScheduledTask
			// 假设 Redis 存储的数据结构是正确的，可以直接映射
			task.ID, _ = strconv.Atoi(strings.TrimPrefix(key, "task:"))
			task.TaskName = taskData["taskName"]
			task.Status = taskData["status"]
			task.NextRunTime, _ = time.Parse(time.RFC3339, taskData["nextRunTime"])
			intervalSeconds, _ := strconv.Atoi(taskData["intervalSeconds"])
			task.IntervalSeconds = &intervalSeconds
			if task.Status == "pending" || task.Status == "running" {
				tasks = append(tasks, task)
			}
		}
		return tasks, nil
	} else {
		var ts []models.ScheduledTask
		err = global.OWEN_DB.Model(&models.ScheduledTask{}).Where("status = ? OR status = ?", "pending", "running").Find(&ts).Error
		if err != nil {
			return takss, err
		}
		return ts, err
	}

}

func (ScheduledTaskService) updateTaskStatus(taskID int, status string) error {
	if global.OWEN_CONFIG.System.TaskDB == "redis" {
		return global.OWEN_REDIS.HSet(global.Ctx, fmt.Sprintf("task:%d", taskID), "status", status).Err()
	} else {
		return global.OWEN_DB.Model(&models.ScheduledTask{}).Where("ID=?", taskID).Updates(map[string]interface{}{
			"status": status,
		}).Error
	}

}
func (ScheduledTaskService) updateTaskInternal(taskID int, nexttime time.Time) error {
	if global.OWEN_CONFIG.System.TaskDB == "redis" {
		return global.OWEN_REDIS.HSet(global.Ctx, fmt.Sprintf("task:%d", taskID), "nextRunTime", nexttime.Format(time.RFC3339), "status", "pending").Err()
	} else {
		return global.OWEN_DB.Model(&models.ScheduledTask{}).Where("ID=?", taskID).Updates(map[string]interface{}{
			"next_run_time": nexttime,
			"status":        "pending",
		}).Error
	}
}
func (ScheduledTaskService) executeTask(task models.ScheduledTask) {
	fmt.Printf("Executing task: %s\n", task.TaskName)
}

// 获取自增 ID
func (ScheduledTaskService) generateTaskID() (int, error) {
	// 使用 Redis 的 INCR 命令生成一个自增的 ID
	id, err := global.OWEN_REDIS.Incr(global.Ctx, "task_id_counter").Result()
	if err != nil {
		return 0, fmt.Errorf("failed to generate task ID: %v", err)
	}
	return int(id), nil
}

// 添加新任务
func (s ScheduledTaskService) AddScheduledTask(req request.AddScheduledTaskInput) (resp response.AddScheduledTaskDto, err error) {
	if global.OWEN_CONFIG.System.TaskDB == "redis" {
		var resp response.AddScheduledTaskDto
		// 生成一个自增的任务 ID
		taskID, err := s.generateTaskID()
		if err != nil {
			return resp, fmt.Errorf("failed to generate task ID: %v", err)
		}

		// 创建任务数据结构
		task := models.ScheduledTask{
			ID:              taskID,
			TaskName:        req.TaskName,
			ScheduleTime:    req.ScheduleTime,
			IntervalSeconds: req.IntervalSeconds,
			NextRunTime:     req.ScheduleTime,
			Status:          "pending",
		}
		// 将任务添加到 Redis 中
		taskKey := fmt.Sprintf("task:%d", taskID)
		err = global.OWEN_REDIS.HMSet(global.Ctx, taskKey, map[string]interface{}{
			"taskName":        task.TaskName,
			"status":          task.Status,
			"scheduleTime":    task.ScheduleTime.Format(time.RFC3339),
			"nextRunTime":     task.NextRunTime.Format(time.RFC3339),
			"intervalSeconds": task.IntervalSeconds,
		}).Err()
		if err != nil {
			return resp, fmt.Errorf("failed to add task to Redis: %v", err)
		}
		fmt.Printf("Task '%s' added successfully\n", task.TaskName)
		return resp, nil
	} else {
		var task models.ScheduledTask
		task.TaskName = req.TaskName
		task.ScheduleTime = req.ScheduleTime
		task.IntervalSeconds = req.IntervalSeconds
		task.NextRunTime = req.ScheduleTime
		task.Status = "pending"
		// 将任务插入数据库
		err = global.OWEN_DB.Create(&task).Error
		if err != nil {
			return resp, fmt.Errorf("failed to add task: %v", err)
		}
		fmt.Printf("Task '%s' added successfully\n", task.TaskName)
		return resp, err
	}
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
			if task.IntervalSeconds != nil && *task.IntervalSeconds > 0 {
				nextRunTime := time.Now().Add(time.Duration(*task.IntervalSeconds) * time.Second)
				if err := s.updateTaskInternal(task.ID, nextRunTime); err != nil {
					log.Printf("Error updating task %d next run time: %v", task.ID, err)
				}
			}
		}
	}

	return nil
}
