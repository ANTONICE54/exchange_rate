package database

import (
	"database/sql"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/viper"
)

func InitDB() *sql.DB {
	conn, err := sql.Open("postgres", viper.GetString("DB_SOURCE"))
	if err != nil {
		log.Fatal(err)

	}
	return conn
}

func RunDBMigration() {
	migration, err := migrate.New(viper.GetString("MIGRATION_PATH"), viper.GetString("DB_SOURCE"))
	if err != nil {
		log.Fatal("cannot create migration:", err)
	}

	if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("failed to run migrate up:", err)
	}

	log.Println("DB migrated successfully")

}
