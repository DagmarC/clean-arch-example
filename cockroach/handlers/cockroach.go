package handlers

import "github.com/labstack/echo/v4"


type CockroachHandler interface {
	DetectCockroach(ctx echo.Context) error
}
