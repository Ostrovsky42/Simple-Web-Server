package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateUserProc(creator UserCreator) func(ctx echo.Context) error {
	return func(ctx echo.Context) error {
		userId, err := creator.CreateUser()
		if err != nil||userId==-1{
			return echo.NewHTTPError((http.StatusBadGateway))
		}
		return echo.NewHTTPError(http.StatusCreated, fmt.Sprintf(`New user with id=%d`, userId))
	}
}
