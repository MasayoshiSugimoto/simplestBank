package domain

import (
	"log"
)

type Bank struct {
	persons      []Person
	accounts     []Account
	transactions []Transaction
}

func (bank *Bank) GetPersonById(id int) Person {
	for _, person := range bank.persons {
		if person.Id() == id {
			return person
		}
	}
	return Person{}
}

func (bank *Bank) GetAccounts(person *Person) []Account {
	accounts := []Account{}
	for _, account := range bank.accounts {
		if account.PersonId() == person.Id() {
			accounts = append(accounts, account)
		}
	}
	return accounts
}

func (bank *Bank) GetBalance(account *Account) int {
	balance := 0
	for _, transaction := range bank.transactions {
		if transaction.To() == account.Id() {
			balance += transaction.Amount()
		} else if transaction.From() == account.Id() {
			balance -= transaction.Amount()
		}
	}
	return balance
}

func (bank *Bank) SaveAccount(account *Account) {
	bank.accounts = append(bank.accounts, *account)
}

func (bank *Bank) SaveTransaction(transaction *Transaction) {
	bank.transactions = append(bank.transactions, *transaction)
}

func (bank *Bank) SavePerson(person *Person) {
	log.Println("Saving client with id:", person.Id())
	bank.persons = append(bank.persons, *person)
}
