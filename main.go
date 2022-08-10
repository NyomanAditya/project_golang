package main

import (
	"project_golang/models"
	"project_golang/routes"
)

func main() {

    db := models.SetupDB()
    db.AutoMigrate(&models.Task{})
     

    r := routes.SetupRoutes(db)
    r.LoadHTMLGlob("views/**/*.html")
    r.Run()
}