package main

import (
	"context"
	"ginbook/handlers"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var recipesHandler *handlers.RecipesHandler

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	client, _ := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err = client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatal(err)
	}
	collection := client.Database(os.Getenv("MONGO_DATABASE")).Collection("recipes")
	recipesHandler = handlers.NewRecipesHandler(ctx, collection)
	log.Println("Connected to MongoDB")
}

func main() {
	router := gin.Default()
	router.POST("/recipes", recipesHandler.NewRecipeHandler)
	router.GET("/recipes", recipesHandler.ListRecipesHandler)
	router.PUT("/recipes/:id", recipesHandler.UpdateRecipeHandler)
	// router.DELETE("/recipes/:id", DeleteRecipeHandler)
	http.ListenAndServe(":8080", router)
}
