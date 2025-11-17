package migrations

import (
	"embed"
	"log"
	"questionAnswer/common/database"

	"github.com/pressly/goose/v3"
)

//go:embed sql/*.sql
var embedMigrations embed.FS

func RunMigrations() {
	db, _ := database.DB.DB()

	goose.SetBaseFS(embedMigrations)
	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatal("Can't set up dialect for migrations: ", err)
	}

	if err := goose.Up(db, "sql"); err != nil {
		log.Fatal("Can't run migrations: ", err)
	}
}