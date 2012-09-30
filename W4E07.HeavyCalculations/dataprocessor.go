package main

import (
	"code.google.com/p/go-uuid/uuid"
	"errors"
	"time"
)

type CalculationTask struct {
	Id     string
	Status string
	Data   string
	Result string
}

func NewTask() *CalculationTask {
	return &CalculationTask{Id: uuid.NewRandom().String()}
}

var tasks = make(map[string]*CalculationTask)

func sendForProcessing(task *CalculationTask) {
	tasks[task.Id] = task
	task.Status = "Sent for processing"

	time.Sleep(time.Second * 5)
	task.Status = "Processing"

	time.Sleep(time.Second * 5)
	reverse := reverse(task.Data)
	task.Result = reverse
	task.Status = "Finished"
}

func reverse(s string) string {
	temp := []byte(s)
	for i := 0; i < len(s)/2; i++ {
		temp[i], temp[len(temp)-(i+1)] = temp[len(temp)-(i+1)], temp[i]
	}
	return string(temp)
}

func getCalculationTask(id string) (*CalculationTask, error) {
	task, exists := tasks[id]
	if !exists {
		return nil, errors.New("No such order: " + id)
	}
	return task, nil
}

func GetTasks() map[string]*CalculationTask {
	return tasks
}
