package usecases

import (
	"time"

	"github.com/DagmarC/clean-arch-example/cockroach/entities"
	"github.com/DagmarC/clean-arch-example/cockroach/models"
	"github.com/DagmarC/clean-arch-example/cockroach/repositories"
)

type cockroachImpl struct {
	cockroachRepository repositories.Cockroach
	cockroachMessaging  repositories.CockroachMessaging
}

func NewCockroachUsecaseImpl(
	cockroachRepository repositories.Cockroach,
	cockroachMessaging repositories.CockroachMessaging,
) CockroachUsecase {
	return &cockroachImpl{
		cockroachRepository: cockroachRepository,
		cockroachMessaging:  cockroachMessaging,
	}
}

func (u *cockroachImpl) Process(in *models.InsertCockroach) error {
	cockroachData := &entities.InsertCockroachDto{
		Amount: in.Amount,
	}
	if err := u.cockroachRepository.InsertCockroachData(cockroachData); err != nil {
		return err
	}

	notification := &entities.CockroachPushNotificationDto{
		Title:        "Cockroach detected!!!",
		Amount:       in.Amount,
		ReportedTime: time.Now().Format("2006-01-02 15:04:05"),
	}
	if err := u.cockroachMessaging.PushNotification(notification); err != nil {
		return err
	}
	
	return nil
}
