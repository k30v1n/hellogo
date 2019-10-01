package controller

import (
	"../model"
	"../service"
	"github.com/gin-gonic/gin"
)

// AddPerson adds a person
func AddPerson(c *gin.Context) {
	person := model.Person{}
	err := c.BindJSON(&person)
	result := service.AddPerson(person)
	if err != nil {
		c.JSON(500, gin.H{"exception": err.Error()})
	} else {
		c.JSON(200, gin.H{"data": result})
	}
}

// GetPerson returns all people
func GetPerson(c *gin.Context) {
	key := c.Params.ByName("key")
	person, found := service.GetPerson(key)

	if found {
		c.JSON(200, gin.H{"data": person})
	} else {
		c.JSON(200, gin.H{"data": nil})
	}
}

// GetPersons gets all persons previous inserted
func GetPersons(c *gin.Context) {
	persons := service.GetPersons()

	c.JSON(200, gin.H{"data": persons})
}
