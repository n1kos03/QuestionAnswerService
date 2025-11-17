package main

import (
	"questionAnswer/common/database"
	"questionAnswer/common/migrations"
)

func main() {
	database.InitDB()

	migrations.RunMigrations()
}