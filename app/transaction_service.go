package app

import (
	"SimplyWebServer/infrastructure/handlers"
	"SimplyWebServer/library"
	"errors"
)

type transactionService struct {
	Storage Storage
}

func NewTransactionService(Storage Storage) *transactionService {
	if Storage==nil{
		panic("Storage is nil")
	}
	return &transactionService{Storage: Storage}
}

func (t *transactionService) AddToBalance(userId int, amount float32)error{

	if err:=t.Storage.AddToBalance(userId, amount);err!=nil{
		return err
	}
	if err:=t.Storage.AddTransaction(library.Adding,userId,amount);err!=nil{
		return err
	}
	return nil
	}

func (t *transactionService) AddTransfer(fromId int,toId int, amount float32)error {
	if !handlers.ValidateParams(amount, fromId, toId){
		return errors.New("400")
	}
	if err:=t.Storage.AddToBalance(fromId,-amount);err!=nil{
		return err
	}
	if err:=t.Storage.AddToBalance(toId,amount);err!=nil{
		return err
	}
	if err:=t.Storage.AddTransaction(library.Transaction,fromId,-amount);err!=nil{
		return err
	}
	if err:=t.Storage.AddTransaction(library.Transaction,toId,amount);err!=nil{
		return err
	}
	return nil
}


