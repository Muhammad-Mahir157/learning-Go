package repositories

import "github.com/Muhammad-Mahir157/clockify-app-clone/domain/entities"

type TimeLogRepository interface {
	Create(timeLogEntity *entities.TimeLog) (*entities.TimeLog, error)
	// GetById(id uint) (*entities.TimeLog, error)
	GetAll() ([]*entities.TimeLog, error)
	// Update(product *entities.TimeLog) (*entities.TimeLog, error)
	// Delete(id uint) error
}
