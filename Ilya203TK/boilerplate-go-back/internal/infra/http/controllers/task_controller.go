package controllers

import "github.com/BohdanBoriak/boilerplate-go-back/internal/app"

type TaskController struct {
	taskService app.TaskService
}

func NewTaskController(ts, app.TaskService) TaskController {
	return TaskController{
		taskService: ts,
	}
}
