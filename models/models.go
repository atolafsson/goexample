package models

import (
	"time"

	"gorm.io/gorm"
)

type Model struct {
	ID          uint `gorm:"primary_key"`
	Created_at  time.Time
	Updated_at  time.Time
	Deleted_att *time.Time
}

type Fact struct {
	gorm.Model
	Question string `json:"question" gorm:"text;not null;default:null"`
	Answer   string `json:"answer" gorm:"text;not null;default:null"`
}
