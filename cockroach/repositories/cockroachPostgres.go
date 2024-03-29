package repositories

import (
	"github.com/DagmarC/clean-arch-example/cockroach/entities"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type cockroachPostgresRepository struct {
	db *gorm.DB
}

func NewCockroachPostgresRepository(db *gorm.DB) Cockroach {
	return &cockroachPostgresRepository{db: db}
}

func (r *cockroachPostgresRepository) InsertCockroachData(in *entities.InsertCockroachDto) error {
	data := &entities.Cockroach{
		Amount: in.Amount,
	}
	result := r.db.Create(data)
	if result != nil {
		log.Errorf("InsertCockroachData: %v", result.Error)
		return result.Error
	}
	log.Debugf("InsertCockroachData: %v", result.RowsAffected)
	return nil
}
