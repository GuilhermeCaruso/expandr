package migrations

import (
	"context"

	"github.com/expandr/expandr/models"
	"github.com/uptrace/bun"
)

func init() {

	Migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
		return db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {

			_, err := tx.NewCreateTable().
				Model((*models.Key)(nil)).
				IfNotExists().
				Exec(ctx)
			if err != nil {
				return err
			}

			_, err = tx.NewCreateTable().
				Model((*models.Collection)(nil)).
				IfNotExists().
				Exec(ctx)
			if err != nil {
				return err
			}

			_, err = tx.NewCreateTable().
				Model((*models.CollectionKey)(nil)).
				IfNotExists().
				Exec(ctx)
			if err != nil {
				return err
			}

			return nil
		})
	}, func(ctx context.Context, db *bun.DB) error {
		return db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
			_, err := tx.NewDropTable().
				Model((*models.CollectionKey)(nil)).
				IfExists().
				Exec(ctx)
			if err != nil {
				return err
			}

			_, err = tx.NewDropTable().
				Model((*models.Collection)(nil)).
				IfExists().
				Exec(ctx)
			if err != nil {
				return err
			}

			_, err = tx.NewDropTable().
				Model((*models.Key)(nil)).
				IfExists().
				Exec(ctx)
			if err != nil {
				return err
			}

			return nil
		})
	})
}
