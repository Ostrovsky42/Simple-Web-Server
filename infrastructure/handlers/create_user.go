package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateUser(creator UserCreator) func(ctx echo.Context) error {
	return func(ctx echo.Context) error {
		userId, err := creator.CreateUser()
		if err != nil {
		}
		return echo.NewHTTPError(http.StatusOK, fmt.Sprintf(`New user with id=%d`, userId))
	}
}
