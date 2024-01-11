package handlers

import "github.com/labstack/echo"

type baseResponse struct {
	Message string `json:"message"`
}

func response(c echo.Context, responseCode int, msg string) error {
	return c.JSON(responseCode, &baseResponse{
		Message: msg,
	})
}