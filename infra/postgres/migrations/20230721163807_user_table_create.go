package migrations

import (
	"database/sql"

	"github.com/pressly/goose/v3"
	model "github.com/vikashkumar2020/quigo-backend/app/models"
	"github.com/vikashkumar2020/quigo-backend/config"
	pgdatabase "github.com/vikashkumar2020/quigo-backend/infra/postgres/database"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	config.LoadEnv()
	config := config.NewDBConfig()
	database := pgdatabase.Database{}
	database.NewDBConnection(config)
	db = database.DB
	goose.AddMigration(upUserTableCreate, downUserTableCreate)
}

func upUserTableCreate(tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	return db.Migrator().CreateTable(&model.User{})
}

func downUserTableCreate(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	return db.Migrator().DropTable(&model.User{})
}
