package app

import (
	"SimplyWebServer/library"
	"github.com/sirupsen/logrus"
)

type transactionService struct {
	Storage Storage
	log     *logrus.Logger
}

func NewTransactionService(Storage Storage, l *logrus.Logger) *transactionService {
	if Storage == nil {
		l.Fatal("Storage is nil")
	}
	return &transactionService{
		Storage: Storage,
		log:     l,
	}
}

func (t *transactionService) AddToBalance(userId int, amount float64) (float64, error) {


	if err:=ValidateParams(amount, userId);err!=nil {
		t.log.Error(err)
		return 0,err
	}

	balance, err := t.Storage.UpdateBalanceByUserId(userId, amount)
	if err != nil {
		t.log.Error("UpdateBalanceByUserId")
		return 0, err
	}
	if err := t.Storage.AddTransaction(library.Adding, userId, amount); err != nil {
		return 0, err
	}
	return balance, nil
}

func (t *transactionService) AddTransfer(fromId int, toId int, amount float64) (float64, float64, error) {

	if err:=ValidateParams(amount, fromId, toId);err!=nil {
		t.log.Error(err)
		return 0,0,err
	}

	fromBalance, err := t.Storage.UpdateBalanceByUserId(fromId, -amount)
	if err != nil {
		return 0, 0, err
	}
	toBalance, err := t.Storage.UpdateBalanceByUserId(toId, amount)
	if err != nil {
		return 0, 0, err
	}
	if err := t.Storage.AddTransaction(library.Transaction, fromId, -amount); err != nil {
		return 0, 0, err
	}
	if err := t.Storage.AddTransaction(library.Transaction, toId, amount); err != nil {
		return 0, 0, err
	}
	return fromBalance, toBalance, nil
}
