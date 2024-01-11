package repositories

import "github.com/DagmarC/clean-arch-example/cockroach/entities"

type CockroachMessaging interface {
	PushNotification(m *entities.CockroachPushNotificationDto) error
}