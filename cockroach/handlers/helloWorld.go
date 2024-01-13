package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Echo defines handler function as func(echo.Context) error where echo.
// Context primarily holds HTTP request and response interfaces.

func HelloWorld(ctx echo.Context) error {
	return response(ctx, http.StatusOK, "HelloWorld! 🪳🪳🪳")
}
