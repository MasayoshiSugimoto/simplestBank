package domain

type Repository struct {
	persons      []Person
	accounts     []Account
	transactions []Transaction
}

func (repository *Repository) GetPersonById(id int) Person {
	for _, person := range repository.persons {
		if person.Id() == id {
			return person
		}
	}
	return Person{}
}

func (repository *Repository) GetAccounts(person *Person) []Account {
	accounts := []Account{}
	for _, account := range repository.accounts {
		if account.PersonId() == person.Id() {
			accounts = append(accounts, account)
		}
	}
	return accounts
}

func (repository *Repository) GetBalance(account *Account) int {
	balance := 0
	for _, transaction := range repository.transactions {
		if transaction.To() == account.Id() {
			balance += transaction.Amount()
		} else if transaction.From() == account.Id() {
			balance -= transaction.Amount()
		}
	}
	return balance
}

func (repository *Repository) SaveAccount(account *Account) {
	repository.accounts = append(repository.accounts, *account)
}

func (repository *Repository) SaveTransaction(transaction *Transaction) {
	repository.transactions = append(repository.transactions, *transaction)
}

func (repository *Repository) SavePerson(person *Person) {
	repository.persons = append(repository.persons, *person)
}
