package usecases

import "github.com/DagmarC/clean-arch-example/cockroach/models"

type CockroachUsecase interface {
	Process(in *models.InsertCockroach) error
}
