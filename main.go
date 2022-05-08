package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Recipe struct {
	Name         string    `json:"name"`
	Tags         []string  `json:"tags"`
	Ingredients  []string  `json:"ingredients"`
	Instructions []string  `json:"instructions"`
	PublishedAt  time.Time `json:"publishedAt"`
}

func IndexHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Working",
	})
}

func main() {
	router := gin.Default()
	router.GET("/", IndexHandler)
	http.ListenAndServe(":8080", router)
}
