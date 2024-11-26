package migrations

import (
	"context"
	"database/sql"

	"github.com/martinmunillas/otter/migrate"
)

func init() {
	migrate.AddMigration("20240909220054_init", upInit, downInit)
}

func upInit(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.ExecContext(ctx, "SELECT 1;")

	return err
}

func downInit(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.ExecContext(ctx, "SELECT 1")
	return err
}

