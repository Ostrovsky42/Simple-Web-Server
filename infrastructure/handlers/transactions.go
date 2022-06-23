package handlers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
)

type transactionHandler struct {
	Service TransactionService
	log     *logrus.Logger
}

func NewTransactionHandler(Service TransactionService, l *logrus.Logger) *transactionHandler {
	if Service == nil {
		l.Fatal("Service is nil")
	}
	return &transactionHandler{
		Service: Service,
		log:     l,
	}
}

func (t *transactionHandler) AddToBalance(ctx echo.Context) error {
	req := new(AddToBalanceRequest)

	err := ctx.Bind(&req)
	if err != nil || !ValidateParams(req.Amount, req.UserId) {
		return echo.NewHTTPError(http.StatusBadRequest, "Wrong input data")
	}

	newBalance, err := t.Service.AddToBalance(req.UserId, req.Amount)
	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	return echo.NewHTTPError(http.StatusOK, fmt.Sprintf(`Now user with id=%d has %.2f on the balance`, req.UserId, newBalance))
}

func (t *transactionHandler) AddTransfer(ctx echo.Context) error {
	req := new(AddTransferRequest)
	err := ctx.Bind(&req)

	if err != nil || !ValidateParams(req.Amount, req.FromId, req.ToId) {
		return echo.NewHTTPError(http.StatusBadRequest, "Wrong input data")
	}

	from, to, err := t.Service.AddTransfer(req.FromId, req.ToId, req.Amount)
	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}
	return echo.NewHTTPError(http.StatusOK, fmt.Sprintf(`Now user with id=%d has %.2f on the balance and user with id=%d has %.2f on the balance`, req.FromId, from, req.ToId, to))
}
