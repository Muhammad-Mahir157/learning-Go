package postgresDb

import (
	"fmt"

	"github.com/Muhammad-Mahir157/clockify-app-clone/domain/entities"
	"github.com/Muhammad-Mahir157/clockify-app-clone/domain/repositories"
	"gorm.io/gorm"
)

type TimeLogRepository struct {
	DB *gorm.DB
}

func NewTimeLogRepository(db *gorm.DB) repositories.TimeLogRepository {
	return &TimeLogRepository{DB: db}
}

func (repo *TimeLogRepository) Create(timeLogEntity *entities.TimeLog) (*entities.TimeLog, error) {

	//mapping the entities into models here ...
	timeLogDbModel := fromEntitytoDbModel(timeLogEntity)
	fmt.Println("DB Model inserted: ", timeLogDbModel)
	err := repo.DB.Create(timeLogDbModel).Error
	if err != nil {
		return nil, err
	}

	//return repo.FindById(dbProduct.Id),nil
	return nil, nil
}

// func GetById(id uint) (*entities.TimeLog, error){

// }

func (repo *TimeLogRepository) GetAll() ([]*entities.TimeLog, error) {
	loggedTimeList := []TimeLog{}

	err := repo.DB.Find(&loggedTimeList).Error
	if err != nil {
		return nil, err
	}
	//mapping the db model into entity model
	entityTimeLogList := make([]*entities.TimeLog, len(loggedTimeList))
	for i, dbTimeLog := range loggedTimeList {
		entityTimeLogList[i] = fromDbToEntityModel(&dbTimeLog)
	}

	return entityTimeLogList, nil
}

// func Update(product *entities.TimeLog) (*entities.TimeLog, error){

// }

// func Delete(id uint) error {

// }
