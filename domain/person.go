package domain

import (
	"fmt"
	"log"
)

var personCounter = 0

type Person struct {
	id        int
	firstName string
	lastName  string
	telephone int
}

func test() {
	fmt.Println("vim-go")
}

func NewPerson(firstName string, lastName string, telephone int) *Person {
	personCounter++
	person := Person{
		id:        personCounter,
		firstName: firstName,
		lastName:  lastName,
		telephone: telephone,
	}
	log.Println("Creating client:", person)
	return &person
}

func (person *Person) Id() int {
	return person.id
}
