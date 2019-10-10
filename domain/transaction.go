package domain

var transactionCounter = 0

type Transaction struct {
	id     int
	from   int
	to     int
	amount int
}

func CreateTransaction(from *Account, to *Account, amount int) *Transaction {
	transactionCounter++
	result := Transaction{
		id:     transactionCounter,
		from:   from.Id(),
		to:     to.Id(),
		amount: amount}
	return &result
}

func (transaction *Transaction) From() int {
	return transaction.from
}

func (transaction *Transaction) To() int {
	return transaction.to
}

func (transaction *Transaction) Amount() int {
	return transaction.amount
}
