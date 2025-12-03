package main

import (
	"errors"
	"flag"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/thelamedev/mattertui/internal/config"
)

func main() {
	var direction string
	flag.StringVar(&direction, "direction", "up", "migration direction (up or down)")
	flag.Parse()

	cfg, err := config.LoadConfig(".", "..")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	dbURL := cfg.Database.Url
	if dbURL == "" {
		log.Fatal("database url is empty")
	}

	// The pgx driver for migrate expects the scheme to be "pgx" or "postgres"
	// If using pgx driver explicitly, we might need to adjust the URL scheme if it's not detected correctly,
	// but usually "postgres://" works if the driver is registered.
	// However, the import _ "github.com/golang-migrate/migrate/v4/database/pgx" registers "pgx" driver.
	// Let's ensure the URL is compatible.
	// Actually, for standard postgres URL, "postgres" driver is usually used (lib/pq).
	// For pgx, we might need to change scheme to "pgx" or rely on auto-detection if "postgres" is also registered.
	// But since we only imported pgx driver, we should probably use "pgx" scheme or check if it handles "postgres".
	// According to docs, it registers "pgx".

	// Let's try to use the URL as is first, but if it fails, we might need to replace "postgres://" with "pgx://"
	// or ensure the driver name in migrate.New is correct.

	m, err := migrate.New(
		"file://migrations",
		dbURL,
	)
	if err != nil {
		log.Fatalf("failed to create migrate instance: %v", err)
	}

	var migErr error
	switch direction {
	case "up":
		fmt.Println("Applying up migrations...")
		migErr = m.Up()
	case "down":
		fmt.Println("Applying down migrations...")
		migErr = m.Down()
	default:
		log.Fatalf("invalid direction: %s", direction)
	}

	if migErr != nil {
		if errors.Is(migErr, migrate.ErrNoChange) {
			fmt.Println("No changes to apply")
		} else {
			log.Fatalf("migration failed: %v", migErr)
		}
	} else {
		fmt.Println("Migration successful")
	}
}
