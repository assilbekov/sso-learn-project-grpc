package migrator

import (
	"errors"
	"flag"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
)

func main() {
	var storagePath, migrationPath, migrationTable string

	flag.StringVar(&storagePath, "storage-path", "", "path to the storage")
	flag.StringVar(&migrationPath, "migration-path", "", "path to the migration files")
	flag.StringVar(&migrationTable, "migration-table", "migrations", "migration table name")
	flag.Parse()

	if storagePath == "" {
		panic("storage path is required")
	}
	if migrationPath == "" {
		panic("migration path is required")
	}

	m, err := migrate.New(
		"file://"+migrationPath,
		fmt.Sprintf("sqlite3://%s?x-migration-table=%s", storagePath, migrationTable))
	if err != nil {
		panic(err)
	}

	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			fmt.Println("no migrations to apply")

			return
		}

		panic(err)
	}

	fmt.Println("migrations applied successfully")
}
