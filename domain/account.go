package domain

var accountCounter = 0

type Account struct {
	id       int
	clientId int
}

func CreateAccount(client *Client) *Account {
	result := Account{
		id:       accountCounter,
		clientId: client.Id(),
	}
	accountCounter++
	return &result
}

func (account *Account) Id() int {
	return account.id
}

func (account *Account) ClientId() int {
	return account.clientId
}
