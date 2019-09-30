package main

import (
	"./controller"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	personGroup := router.Group("person/v1")
	{
		personGroup.POST("/person", controller.AddPerson)
		personGroup.GET("/person", controller.GetPersons)
		personGroup.GET("/person/:index", controller.GetPerson)
	}
	router.Use(LoggingMiddleware)
	router.Run(":8080")

	// go build .
}

// LoggingMiddleware to print all the requests
func LoggingMiddleware(c *gin.Context) {

	reqElements := make(map[string]interface{})
	reqElements["path"] = c.Request.URL.Path
	println(reqElements)
	c.Next()
}
