package repositories

import (
	"github.com/Muhammad-Mahir157/clockify-app-clone/domain/entities"
	"github.com/google/uuid"
)

type TimeLogRepository interface {
	Create(timeLogEntity *entities.TimeLog) (*entities.TimeLog, error)
	GetById(id uuid.UUID) (*entities.TimeLog, error)
	GetAll() ([]*entities.TimeLog, error)
	Update(product *entities.TimeLog) (*entities.TimeLog, error)
	Delete(id uuid.UUID) (bool, error)
}
