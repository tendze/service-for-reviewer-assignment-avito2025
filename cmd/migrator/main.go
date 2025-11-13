package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"

	"dang.z.v.task/internal/config"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	_ "github.com/lib/pq"
)

func main() {
	configPath := flag.String("config", "", "Path to config file")
	migrationsDir := flag.String("dir", "./migrations", "Migrations directory")
	direction := flag.String("direction", "up", "Migration direction: up or down")
	flag.Parse()

	if *configPath != "" {
		os.Setenv("CONFIG_PATH", *configPath)
	}

	cfg := config.MustLoad()

	dsn := cfg.DB.DSN()
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("failed to open DB: %v", err)
	}
	defer db.Close()

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatalf("failed to create DB driver: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://"+*migrationsDir,
		"postgres",
		driver,
	)
	if err != nil {
		log.Fatalf("failed to init migrate: %v", err)
	}

	switch *direction {
	case "up":
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatalf("failed to apply migrations: %v", err)
		}
		fmt.Println("Migrations applied successfully")
	case "down":
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatalf("failed to rollback migrations: %v", err)
		}
		fmt.Println("Migrations rolled back successfully")
	default:
		fmt.Println("unknown direction, use -dir=up or -dir=down")
		os.Exit(1)
	}
}
