package main

import "fmt"
import "domain"

var (
	repository domain.Repository
)

func main() {
	john := domain.CreatePerson("john", "carmack", 123456)
	repository.SavePerson(john)
	johnAccount := domain.CreateAccount(john)
	repository.SaveAccount(johnAccount)

	bob := domain.CreatePerson("bob", "martin", 345678)
	repository.SavePerson(bob)
	bobAccount := domain.CreateAccount(bob)
	repository.SaveAccount(bobAccount)

	transaction := domain.CreateTransaction(johnAccount, bobAccount, 100)
	repository.SaveTransaction(transaction)

	fmt.Println(repository.GetBalance(johnAccount))
	fmt.Println(repository.GetBalance(bobAccount))
}
