package postgresDb

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TimeLog struct {
	TimeLogId     uuid.UUID `gorm:"primary-key;autoIncrement;column:time_log_id"`
	StartTime     time.Time `gorm:"column:start_time"`
	EndTime       time.Time `gorm:"column:end_time"`
	LoggedAt      time.Time `gorm:"column:logged_at"`
	TotalDuration string    `gorm:"column:total_duration"`
	Description   string    `gorm:"column:description"`
}

func MigrateTimelog(db *gorm.DB) error {
	err := db.AutoMigrate(&TimeLog{})
	return err
}
