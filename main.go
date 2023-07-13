package main

import (
	"go-htmx/controller" 
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("C:/Users/JASIM MANSOOR/go/src/go-htmx/index.html")


	// Define your routes here
	router.GET("/", controller.Hometemp)
	router.POST("/add-club/",controller.Addtemp)


	router.Run("localhost:8088")
}

