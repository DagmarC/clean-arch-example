package handlers

import "github.com/labstack/echo"

type Cockroachhandler interface {
	DetectCockroach(ctx echo.Context) error
}