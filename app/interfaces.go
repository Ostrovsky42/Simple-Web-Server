package app

import (
	"SimplyWebServer/library"
)

//go:generate mockgen -source=interfaces.go -destination=mocks.go -package=app
type Storage interface {
	CreateUser() (int, error)
	UpdateBalanceByUserId(userId int, amount float64) (float64, error)
	AddTransaction(transactType library.TransactionType, userId int, amount float64) error
}