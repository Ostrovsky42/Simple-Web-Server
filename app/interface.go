package app

import(
	"SimplyWebServer/library"
)

type Storage interface{
	CreateUser()(int, error)
	AddToBalance(userId int, amount float32)error
	AddTransaction(transactType library.TransactionType, userId int, amount float32)error
	}	