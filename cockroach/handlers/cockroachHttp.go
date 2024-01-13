package handlers

import (
	"net/http"

	"github.com/DagmarC/clean-arch-example/cockroach/models"
	"github.com/DagmarC/clean-arch-example/cockroach/usecases"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type cockroachHttpHandler struct {
	cockroach usecases.CockroachUsecase
}

func NewCockroachHttpHandler(cockroach usecases.CockroachUsecase) CockroachHandler {
	return &cockroachHttpHandler{cockroach: cockroach}
}

func (h *cockroachHttpHandler) DetectCockroach(ctx echo.Context) error {
	req := new(models.InsertCockroach)
	if err := ctx.Bind(req); err != nil {
		log.Errorf("Error binding request body: %v", req)
		return response(ctx, http.StatusBadRequest, "Bad request")
	}
	// Check if the NumberOfCockroaches field is valid (you may add more validation)
	if req.Amount <= 0 {
		return response(ctx, http.StatusBadRequest, "Number of cockroaches must be greater than zero")
	}

	if err := h.cockroach.Process(req); err != nil {
		return response(ctx, http.StatusInternalServerError, "Processing data failed")
	}

	return response(ctx, http.StatusOK, "Success 🪳🪳🪳")
}
