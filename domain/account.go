package domain

var accountCounter = 0

type Account struct {
	id       int
	personId int
}

func CreateAccount(person *Person) *Account {
	result := Account{
		id:       accountCounter,
		personId: person.Id(),
	}
	accountCounter++
	return &result
}

func (account *Account) Id() int {
	return account.id
}

func (account *Account) PersonId() int {
	return account.personId
}
