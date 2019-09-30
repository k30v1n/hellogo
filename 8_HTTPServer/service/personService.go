package service

import (
	model "../model"
)

var memoryData = []model.Person{}

// AddPerson adds a new person
func AddPerson(person model.Person) {
	memoryData = append(memoryData, person)
}

// GetPerson returns a person from a specific index
func GetPerson(index int64) *model.Person {
	if index > 0 && index <= int64(len(memoryData)) {
		return &memoryData[index-1]
	}
	return nil
}

// GetPersons returns all persons
func GetPersons() []model.Person {
	return memoryData
}
