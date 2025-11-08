package main

import (
    "log"

    "github.com/gofiber/fiber/v2"
    
)
import (
    "context"
    "time"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var quizCollection *mongo.Collection

func main() {
    app := fiber.New()

    app.Get("/", index)
	app.Get("/api/quizzes", getQuizzes)

	log.Fatal(app.Listen(":3000"))
}

func setupDb(){
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err) //crash the app
	}

	quizCollection := client.Database("quiz").Collection("quizzes")
}

func index(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func getQuizzes(c *fiber.Ctx) error {
	cursor, err := quizCollection.Find(context.Background(), bson.M{})
	if err != nil {
		panic(err)
	}
	list := []map[string]any{
		map[string]any{
			"id":    1,
			"title": "Sample Quiz",
		},
	}
	return c.JSON(list)
}