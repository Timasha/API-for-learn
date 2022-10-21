package main

import (
	"api-for-learn/internal/api"
	"api-for-learn/internal/cases"
	"api-for-learn/internal/storage"
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {

	var mongo storage.Mongo
	connErr := mongo.Connect(context.Background(), "0.0.0.0", "27017")

	if connErr != nil {
		log.Fatalf("mongo connect error: %v", connErr)
	}
	defer mongo.Close()

	var userCase cases.UserCase
	userCase.New(&mongo)

	app := fiber.New()
	api.CreateUserApi(context.Background(), &userCase, app)
	app.Listen(":8080")
}
