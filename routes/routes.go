package routes

import (
	"project_golang/controllers"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})
	r.GET("/", controllers.Root)
	r.GET("/tasks", controllers.FindTasks)
	r.POST("/tasks", controllers.CreateTask)
	r.GET("/tasks/:id", controllers.FindTask)
	r.PATCH("/tasks/:id", controllers.UpdateTask)
	r.POST("/tasks/update/:id", controllers.UpdateTask)
	r.DELETE("tasks/:id", controllers.DeleteTask)
	r.GET("tasks/delete/:id", controllers.DeleteTask)
	r.GET("/tasks/:id/done", controllers.DoneTask)
	r.GET("/tasks/tambah", controllers.Tambah)
	return r
}