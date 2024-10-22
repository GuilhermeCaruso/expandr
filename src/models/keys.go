package models

import (
	"github.com/google/uuid"
)

type Key struct {
	Metadata
	KeyName string `json:"key_name"`
	Value   string `json:"value"`
}

type Collection struct {
	Metadata
	Name string `json:"name"`
	Keys []Key  `gorm:"many2many:collection_keys"`
}

type CollectionKey struct {
	CollectionID uuid.UUID `gorm:"primaryKey"`
	Collection   *Collection
	KeyID        uuid.UUID `gorm:"primaryKey"`
	Key          *Key
}
