package models

import (
	"github.com/google/uuid"
)

type Key struct {
	Metadata
	KeyName string
	Value   string
}

type Collection struct {
	Metadata
	Name string
	Keys []Key
}

type CollectionKey struct {
	CollectionID uuid.UUID
	Collection   *Collection
	KeyID        uuid.UUID
	Key          *Key
}
