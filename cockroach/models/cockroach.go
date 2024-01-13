package models

// model that will be used to receive the request for cockroach data recording

type InsertCockroach struct {
	Amount uint32 `json:"amount"`
}
