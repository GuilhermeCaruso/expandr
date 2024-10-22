package migrations

import (
	"context"

	"github.com/expandr/expandr/pkg/database"
	"gorm.io/gorm"
)

func init() {
	Migrator.RegisterMigration(
		database.Migration{
			Name: "initial2",
			Up: func(ctx context.Context, db *gorm.DB) error {
				return nil
			},
			Down: func(ctx context.Context, db *gorm.DB) error {
				return nil
			},
		},
	)
}
