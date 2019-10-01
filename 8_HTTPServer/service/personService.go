package service

import (
	dal "../dal"
	model "../model"
)

// AddPerson adds a new person
func AddPerson(person model.Person) string {
	return dal.SetPerson(person)
}

// GetPerson returns a person from a specific index
func GetPerson(key string) (dal.PersonDal, bool) {
	return dal.GetPerson(key)
}

// GetPersons returns all persons
func GetPersons() []dal.PersonDal {
	return dal.GetPersons()
}
