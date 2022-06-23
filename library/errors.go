package library

import "github.com/pkg/errors"

var (
ErrSubZeroAmount =errors.New("amount mast be greater than zero")
ErrSubZeroID     =errors.New("id mast be greater than zero")
ErrEqualIDs      =errors.New("it is not possible to transfer to your account")
ErrWrongData 	 =errors.New("wrong input data")
ErrBalanceBeZero =errors.New("can`t perform the operation, the user balance will be below zero")
)
