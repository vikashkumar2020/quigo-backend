package migrations

import (
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
	model "github.com/vikashkumar2020/quigo-backend/app/models"
)

func init() {
	goose.AddMigrationContext(upRideTable, downRideTable)
}

func upRideTable(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	return db.Migrator().CreateTable(&model.Rides{})
}

func downRideTable(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	return db.Migrator().DropTable(&model.Rides{})
}
