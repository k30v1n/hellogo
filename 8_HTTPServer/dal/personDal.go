package dal

import (
	"../model"
	"github.com/google/uuid"
	"github.com/patrickmn/go-cache"
)

var c *cache.Cache

// PersonDal is the object saved
type PersonDal struct {
	ID   string
	Name string
	Age  string
}

func initializeCache() {
	if c == nil {
		c = cache.New(0, 0)
	}
}

// SetPerson saves a person
func SetPerson(person model.Person) string {
	initializeCache()

	uuid := uuid.New().String()

	personDal := PersonDal{ID: uuid, Name: person.Name, Age: person.Age}
	c.Set(uuid, personDal, cache.DefaultExpiration)

	return uuid
}

// GetPerson gets a person using a key
func GetPerson(key string) (PersonDal, bool) {
	initializeCache()

	item, found := c.Get(key)
	if found {
		return item.(PersonDal), found
	}
	return PersonDal{}, found
}

// GetPersons returns all saved items
func GetPersons() []PersonDal {
	initializeCache()

	items := c.Items()
	var result = []PersonDal{}
	for _, value := range items {
		result = append(result, value.Object.(PersonDal))
	}
	return result
}
