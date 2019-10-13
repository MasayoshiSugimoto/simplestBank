package domain

import (
	"log"
)

type Bank struct {
	clients      []Client
	accounts     []Account
	transactions []Transaction
}

func (bank *Bank) GetclientById(id int) Client {
	for _, client := range bank.clients {
		if client.Id() == id {
			return client
		}
	}
	return Client{}
}

func (bank *Bank) GetAccounts(client *Client) []Account {
	accounts := []Account{}
	for _, account := range bank.accounts {
		if account.ClientId() == client.Id() {
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

func (bank *Bank) SaveClient(client *Client) {
	log.Println("Saving client with id:", client.Id())
	bank.clients = append(bank.clients, *client)
}
