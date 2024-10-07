package models

import (
	"time"

	"gorm.io/gorm"
)

type Schedule struct {
	gorm.Model
	Name      string    `gorm:"type:text;not null" json:"name"` // Name of the schedule
	StartTime time.Time `gorm:"not null" json:"start_time"`     // Starting time of the schedule
	EndTime   time.Time `gorm:"not null" json:"end_time"`       // Ending time of the schedule
}
