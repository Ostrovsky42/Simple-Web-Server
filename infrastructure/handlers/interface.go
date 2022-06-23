package handlers

type TransactionService interface {
	AddToBalance(userId int, amount float32) (float32, error)
	AddTransfer(fromId int, toId int, amount float32) (float32, float32, error)
}

type UserCreator interface {
	CreateUser() (int, error)
}
