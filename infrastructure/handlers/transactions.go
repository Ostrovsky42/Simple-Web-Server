package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type transactionHandler struct {
	Service TransactionService
}

func NewTransactionHandler(Service TransactionService) *transactionHandler {
	if Service == nil {
		panic("Service is nil")
	}
	return &transactionHandler{Service: Service}
}

func (t *transactionHandler) AddToBalance(ctx echo.Context) error {
	req := new(AddToBalanceRequest)
	err := ctx.Bind(&req)

	if err != nil||!ValidateParams(req.Amount, req.UserId){
		return echo.NewHTTPError(http.StatusBadRequest, "Wrong input data")
	}

	if err := t.Service.AddToBalance(req.UserId, req.Amount); err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	return nil
}

func (t *transactionHandler) AddTransfer(ctx echo.Context) error {
	req := new(AddTransferRequest)
	err := ctx.Bind(&req)

	if err != nil||!ValidateParams(req.Amount, req.FromId,req.ToId) {
		return echo.NewHTTPError(http.StatusBadRequest, "Wrong input data")
	}

	if err := t.Service.AddTransfer(req.FromId, req.ToId, req.Amount); err != nil {
		return echo.NewHTTPError(400, err.Error())
	}
	return nil
}
