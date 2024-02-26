package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/nishanth-gowda/tasker/db"
	"github.com/nishanth-gowda/tasker/handlers"
	"github.com/nishanth-gowda/tasker/model"
	"github.com/nishanth-gowda/tasker/repository"
	"github.com/nishanth-gowda/tasker/routes"
)

func init() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal("Error getting current working directory: ", err)
	}
	log.Println("Current working directory:", cwd)

	_, err = os.Stat(".env")
	if err != nil {
		log.Println("Error: .env file not found")
	}

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file: ", err)
	}

}
func main() {
	db, err := db.GetConnection()
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&model.Project{})
	db.AutoMigrate(&model.Task{})


	projectRepo := repository.NewProjectRepository(db)
	taskRepo := repository.NewTaskRepository(db)

	projectHandler := handlers.NewProjectHandler(projectRepo)
	taskHandler := handlers.NewTaskHandler(taskRepo)



	r := gin.Default()

	routes.RegisterRoutes(r, projectHandler, taskHandler)


	r.Run(":8000")

}
