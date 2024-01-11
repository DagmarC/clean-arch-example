package repositories

import (
	"github.com/DagmarC/clean-arch-example/cockroach/entities"
	"github.com/labstack/gommon/log"
)

type cockroachFCMMessaging struct{}

func NewCockroachFCMMessaging() CockroachMessaging {
	return &cockroachFCMMessaging{}
}

func (cm *cockroachFCMMessaging) PushNotification(m *entities.CockroachPushNotificationDto) error {
	// ... handle logic to push FCM notification here
	log.Debugf("Pushed FCM notification with data %v", m)
	return nil
}