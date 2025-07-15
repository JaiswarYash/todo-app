package main

import (
	"context"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"todo-app/backend/config"
	"todo-app/backend/controllers"
	"todo-app/backend/middlewares"
	"todo-app/backend/routes"
	"todo-app/backend/services"
)

func main() {
	cfg := config.LoadConfig()

	port := cfg.Port
	if port == "" {
		port = "8080"
	}
	mongoURI := cfg.MongoURI
	if mongoURI == "" {
		mongoURI = "mongodb://localhost:27017"
	}
	mongoDBName := cfg.MongoDB
	if mongoDBName == "" {
		mongoDBName = "todoapp"
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatalf("failed to connect to MongoDB: %v", err)
	}
	db := client.Database(mongoDBName)
	tasksCollection := db.Collection("tasks")

	r := gin.Default()
	r.Use(middlewares.CORSMiddleware())

	taskService := services.NewTaskService(tasksCollection)
	taskController := controllers.NewTaskController(taskService)

	routes.RegisterTaskRoutes(r, taskController)

	log.Printf("Server running on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
