package controller

import (
	"strconv"

	"../model"
	"../service"
	"github.com/gin-gonic/gin"
)

// AddPerson adds a person
func AddPerson(c *gin.Context) {
	person := model.Person{}
	err := c.BindJSON(&person)
	service.AddPerson(person)
	exception := ""
	if err != nil {
		exception = err.Error()
		c.JSON(500, gin.H{"exception": exception})
	} else {
		c.JSON(200, gin.H{"data": person})
	}
}

// GetPerson returns all people
func GetPerson(c *gin.Context) {
	index, _ := strconv.ParseInt(c.Params.ByName("index"), 10, 64)
	person := service.GetPerson(index)

	c.JSON(200, gin.H{"data": person})
}

// GetPersons gets all persons previous inserted
func GetPersons(c *gin.Context) {
	persons := service.GetPersons()

	c.JSON(200, gin.H{"data": persons})
}
