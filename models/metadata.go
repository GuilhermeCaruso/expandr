package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Metadata struct {
	ID        uuid.UUID    `bun:"id,type:uuid,default:gen_random_uuid(),pk"`
	CreatedAt time.Time    `bun:"created_at,nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time    `bun:"updated_at,nullzero,notnull,default:current_timestamp"`
	DeletedAt bun.NullTime `bun:"deleted_at"`
}
