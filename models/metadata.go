package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Metadata struct {
	ID        uuid.UUID      `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at,omitempty" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at,omitempty" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}

func (r *Metadata) BeforeCreate(tx *gorm.DB) (err error) {
	if r.ID == uuid.Nil {
		r.ID = uuid.New()
	}
	return
}
