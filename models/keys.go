package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Key struct {
	bun.BaseModel `bun:"table:keys,alias:k"`
	Metadata
	KeyName string `bun:"key_name,type:varchar,notnull"`
	Value   string `bun:"value,type:varchar"`
}

type Collection struct {
	bun.BaseModel `bun:"table:collections,alias:c"`
	Metadata
	Name string `bun:"name,type:varchar,notnull"`
	Keys []Key  `bun:"m2m:order_to_items,join:Collection=Key"`
}

type CollectionKey struct {
	CollectionID uuid.UUID   `bun:"collection_id,pk"`
	Collection   *Collection `bun:"rel:belong-to,join:order_id=id"`
	KeyID        uuid.UUID   `bun:"key_id,pk"`
	Key          *Key        `bun:"rel:belong-to,join:order_id=id"`
}
