package postgresDb

import (
	"fmt"

	"github.com/Muhammad-Mahir157/clockify-app-clone/domain/entities"
	"github.com/Muhammad-Mahir157/clockify-app-clone/domain/repositories"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TimeLogRepository struct {
	DB *gorm.DB
}

func NewTimeLogRepository(db *gorm.DB) repositories.TimeLog {
	return &TimeLogRepository{DB: db}
}

func (repo *TimeLogRepository) Create(timeLogEntity *entities.TimeLog) (*entities.TimeLog, error) {

	//mapping the entity into db model
	timeLog := fromEntitytoDbModel(timeLogEntity)

	err := repo.DB.Create(timeLog).Error
	if err != nil {
		return nil, err
	}

	return repo.GetById(timeLog.TimeLogId)
	//return nil, nil
}

func (repo *TimeLogRepository) GetById(timeLogId uuid.UUID) (*entities.TimeLog, error) {
	timeLog := &TimeLog{}

	err := repo.DB.Where("time_log_id = ?", timeLogId).First(timeLog).Error
	if err != nil {
		return nil, err
	}

	//mapping the db model into entity model
	return fromDbToEntityModel(timeLog), nil
}

func (repo *TimeLogRepository) GetAll() ([]*entities.TimeLog, error) {
	timeLogList := []TimeLog{}

	err := repo.DB.Find(&timeLogList).Error
	if err != nil {
		return nil, err
	}

	//mapping the db model into entity model
	entityTimeLogList := make([]*entities.TimeLog, len(timeLogList))
	for i, dbTimeLog := range timeLogList {
		entityTimeLogList[i] = fromDbToEntityModel(&dbTimeLog)
	}

	return entityTimeLogList, nil
}

func (repo *TimeLogRepository) Update(updatedTimeLogEntity *entities.TimeLog) (*entities.TimeLog, error) {

	//mapping the entity into db model
	timeLog := fromEntitytoDbModel(updatedTimeLogEntity)

	err := repo.DB.Where("time_log_id = ?", timeLog.TimeLogId).Save(timeLog).Error
	if err != nil {
		return nil, err
	}

	return repo.GetById(timeLog.TimeLogId)
}

func (repo *TimeLogRepository) Delete(timeLogId uuid.UUID) (bool, error) {
	timeLog := &TimeLog{}

	err := repo.DB.Delete(timeLog, timeLogId).Error
	if err != nil {
		return false, err
	} else if repo.DB.RowsAffected < 1 {
		return false, fmt.Errorf("row with id=%s cannot be deleted because it doesn't exist", timeLogId)
	}

	return true, nil
}
