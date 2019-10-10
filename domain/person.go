package domain

import "fmt"

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

func CreatePerson(firstName string, lastName string, telephone int) *Person {
	personCounter++
	person := Person{
		id:        personCounter,
		firstName: firstName,
		lastName:  lastName,
		telephone: telephone,
	}
	return &person
}

func (person *Person) Id() int {
	return person.id
}
