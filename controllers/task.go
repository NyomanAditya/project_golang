package controllers

import (
	"net/http"
	"project_golang/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type CreateTaskInput struct {
	AssingedTo string `json:"assingedTo"`
	Task       string `json:"task"`
	Deadline   string `json:"deadline"`
}

type UpdateTaskInput struct {
	AssingedTo string `json:"assingedTo"`
	Task       string `json:"task"`
	Deadline   string `json:"deadline"`
}

type DoneTaskInput struct {
	IsDone 	   bool   `json:"isDone"`

}



// GET /tasks
// Get all tasks
func FindTasks(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var tasks []models.Task
	db.Find(&tasks)

	// c.JSON(http.StatusOK, gin.H{"data": tasks})
	c.HTML(http.StatusOK, "tasks/index", gin.H{
		"title ": "Task Index",
		"tasks" : tasks,
	})
}

// POST /tasks
// Create new task
func CreateTask(c *gin.Context) {
	// Validate input
	var input CreateTaskInput
	// if err := c.ShouldBindJSON(&input); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }
	

	// Create task
	task := models.Task{AssingedTo: input.AssingedTo, Task: input.Task, Deadline: input.Deadline}

	db := c.MustGet("db").(*gorm.DB)
	db.Create(&task)

	// c.JSON(http.StatusOK, gin.H{"data": task})
	c.Redirect(http.StatusFound, "/tasks")
}

// GET /tasks/:id
// Find a task
func FindTask(c *gin.Context) { // Get model if exist
	var task models.Task

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// c.JSON(http.StatusOK, gin.H{"data": task})
	c.HTML(http.StatusOK, "tasks/edit", gin.H{
		"title ": "Task Index",
		"tasks" : task,
	})
}

// PATCH /tasks/:id
// Update a task
func UpdateTask(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var task models.Task
	if err := db.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input UpdateTaskInput
	// if err := c.ShouldBindJSON(&input); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	

	var updatedInput models.Task
	updatedInput.Deadline = input.Deadline
	updatedInput.AssingedTo = input.AssingedTo
	updatedInput.Task = input.Task

	db.Model(&task).Updates(updatedInput)

	// c.JSON(http.StatusOK, gin.H{"data": task})
	c.Redirect(http.StatusFound, "/tasks")
}


// / GET /tasks/:id/done
// Update a task done
func DoneTask(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var task models.Task
	if err := db.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}


	var updatedInput models.Task
	updatedInput.IsDone = !updatedInput.IsDone

	db.Model(&task).Updates(updatedInput)

	// c.JSON(http.StatusOK, gin.H{"data": task})
	c.Redirect(http.StatusFound, "/tasks")
}

// DELETE /tasks/:id
// Delete a task
func DeleteTask(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var book models.Task
	if err := db.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&book)

	// c.JSON(http.StatusOK, gin.H{"data": true})
	c.Redirect(http.StatusFound, "/tasks")
}

func Root(c *gin.Context) {
	c.Redirect(http.StatusFound, "/tasks")
}

func Tambah(c *gin.Context) {
	c.Redirect(http.StatusFound, "/tasks/tambah")
}

