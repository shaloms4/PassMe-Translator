package main

import (
	"context"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/shaloms4/Pass-Me-Core-Functionality/delivery/controllers"
	"github.com/shaloms4/Pass-Me-Core-Functionality/delivery/routers"
	repositories "github.com/shaloms4/Pass-Me-Core-Functionality/repositories"
	usecases "github.com/shaloms4/Pass-Me-Core-Functionality/usecases"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Only load .env if running locally
	if os.Getenv("RENDER") == "" {
		err := godotenv.Load("./.env")
		if err != nil {
			log.Fatalf("Error loading .env file: %v", err)
		}
	}

	// Get MongoDB URI from the environment variables
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		log.Fatal("MONGO_URI is not set in the .env file")
	}

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	// Select the database
	db := client.Database("passme")

	// Initialize repositories
	flightRepo := repositories.NewFlightRepository(db)
	userRepo := repositories.NewUserRepository(db)

	// Initialize use cases
	flightUC := usecases.NewFlightUseCase(flightRepo)
	userUC := usecases.NewUserUseCase(userRepo)

	// Initialize controllers
	flightController := controllers.NewFlightController(flightUC)
	userController := controllers.NewUserController(userUC)

	// Set up the Gin router
	r := gin.Default()

	// Set up the routes
	// Each router file handles its own route setup
	routers.SetupUserRoutes(r, userController)
	routers.SetupFlightRoutes(r, flightController)

	// Start the server
	log.Println("Server is running at :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
