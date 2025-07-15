package controllers

import (
	"net/http"
	"todo-app/backend/services"
	"todo-app/backend/utils"

	"github.com/gin-gonic/gin"
)

type TaskController struct {
	Service *services.TaskService
}

func NewTaskController(service *services.TaskService) *TaskController {
	return &TaskController{Service: service}
}

func (tc *TaskController) CreateTask(c *gin.Context) {
	var input struct {
		Title string `json:"title"`
	}
	if err := c.ShouldBindJSON(&input); err != nil || input.Title == "" {
		utils.JSONError(c, http.StatusBadRequest, "Title is required")
		return
	}
	task, err := tc.Service.Create(input.Title)
	if err != nil {
		utils.JSONError(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated, task)
}

func (tc *TaskController) GetTasks(c *gin.Context) {
	tasks, err := tc.Service.List()
	if err != nil {
		utils.JSONError(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, tasks)
}

func (tc *TaskController) UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var input struct {
		Title     string `json:"title"`
		Completed *bool  `json:"completed"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.JSONError(c, http.StatusBadRequest, "Invalid input")
		return
	}
	task, err := tc.Service.Update(id, input.Title, input.Completed)
	if err != nil {
		if err.Error() == "not found" {
			utils.JSONError(c, http.StatusNotFound, "Task not found")
		} else if err.Error() == "invalid id" {
			utils.JSONError(c, http.StatusBadRequest, "Invalid ID")
		} else {
			utils.JSONError(c, http.StatusInternalServerError, err.Error())
		}
		return
	}
	c.JSON(http.StatusOK, task)
}

func (tc *TaskController) DeleteTask(c *gin.Context) {
	id := c.Param("id")
	err := tc.Service.Delete(id)
	if err != nil {
		if err.Error() == "not found" {
			utils.JSONError(c, http.StatusNotFound, "Task not found")
		} else if err.Error() == "invalid id" {
			utils.JSONError(c, http.StatusBadRequest, "Invalid ID")
		} else {
			utils.JSONError(c, http.StatusInternalServerError, err.Error())
		}
		return
	}
	c.Status(http.StatusNoContent)
}
