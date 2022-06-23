package handlers

type TransactionService interface {
	AddToBalance(userId int, amount float64) (float64, error)
	AddTransfer(fromId int, toId int, amount float64) (float64, float64, error)
}

type UserCreator interface {
	CreateUser() (int, error)
}
