package handlers

import (
	"SimplyWebServer/library"
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
	if err != nil  {
		return echo.NewHTTPError(http.StatusBadRequest, library.ErrWrongData)
	}

	newBalance, err := t.Service.AddToBalance(req.UserId, req.Amount)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return echo.NewHTTPError(http.StatusOK,
		fmt.Sprintf(`Now user with id=%d has %.2f on the balance`,
			req.UserId, newBalance))
}

func (t *transactionHandler) AddTransfer(ctx echo.Context) error {
	req := new(AddTransferRequest)
	err := ctx.Bind(&req)

	if err != nil  {
		return echo.NewHTTPError(http.StatusBadRequest, library.ErrWrongData)
	}

	from, to, err := t.Service.AddTransfer(req.FromId, req.ToId, req.Amount)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return echo.NewHTTPError(http.StatusOK,
		fmt.Sprintf(`Now user with id=%d has %.2f on the balance and user with id=%d has %.2f on the balance`,
			req.FromId, from, req.ToId, to))
}
