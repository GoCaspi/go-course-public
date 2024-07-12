package main

import (
	"errors"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mongodb"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Loading .env failed:", err)
	}

	m, err := migrate.New(os.Getenv("MIGRATION_PATH"), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("loading migration failed:", err)
	}
	err = m.Up()
	if err != nil && errors.Is(err, migrate.ErrNoChange) {
		log.Fatal("executing migration failed:", err)
	}
	log.Print("DB migration success")
}
