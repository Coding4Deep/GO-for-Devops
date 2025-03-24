package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Task struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

var tasks []Task
var nextID = 1

func getTasks(c *gin.Context) {
	c.JSON(http.StatusOK, tasks)
}

func addTask(c *gin.Context) {
	var newTask Task
	if err := c.BindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newTask.ID = nextID
	nextID++
	tasks = append(tasks, newTask)
	c.JSON(http.StatusCreated, newTask)
}

func updateTask(c *gin.Context) {
	var updatedTask Task
	if err := c.BindJSON(&updatedTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, task := range tasks {
		if task.ID == updatedTask.ID {
			tasks[i] = updatedTask
			c.JSON(http.StatusOK, updatedTask)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
}

func deleteTask(c *gin.Context) {
	var taskID int
	if err := c.BindJSON(&taskID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, task := range tasks {
		if task.ID == taskID {
			tasks = append(tasks[:i], tasks[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
}

func main() {
	r := gin.Default()

	r.GET("/tasks", getTasks)
	r.POST("/tasks", addTask)
	r.PUT("/tasks", updateTask)
	r.DELETE("/tasks", deleteTask)

	r.Run(":8080")
}
