package dbmigrate

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// Run applies migrations from migrationsFolder to database.
func Run(database *sql.DB, migrationsFolder string) error {
	// Initialize migrations table, if it does not exist yet
	_, err := database.Exec("create table if not exists migrations(id serial, name text not null, created_at timestamp with time zone not null)")
	if err != nil {
		return err
	}
	_, err = database.Exec("create unique index idx_migrations_name on migrations(name)")
	if err != nil {
		if !strings.Contains(err.Error(), "already exists") {
			return err
		}
	}

	// Scan migration file names in migrations folder
	d, err := os.Open(migrationsFolder)
	if err != nil {
		return err
	}
	dir, err := d.Readdir(-1)
	if err != nil {
		return err
	}

	// Run migrations
	tx, err := database.Begin()
	if err != nil {
		return err
	}
	sqlFiles := make([]string, 0)
	for _, f := range dir {
		if filepath.Ext(f.Name()) == ".sql" {
			sqlFiles = append(sqlFiles, f.Name())
		}
	}
	sort.Strings(sqlFiles)
	for _, filename := range sqlFiles {
		// if exists in migrations table, leave it
		// else execute sql
		var count int
		err := tx.QueryRow("select count(1) from migrations where name = $1", filename).Scan(&count)
		if err != nil {
			return err
		}
		if count > 0 {
			continue // already migrated
		}
		b, err := ioutil.ReadFile(filepath.Join(migrationsFolder, filename))
		if err != nil {
			return err
		}
		migration := string(b)
		if len(migration) == 0 {
			continue // empty file
		}
		_, err = tx.Exec(migration)
		if err != nil {
			return err
		}
		_, err = tx.Exec("insert into migrations(name, created_at) values($1, current_timestamp)", filename)
		if err != nil {
			return err
		}
		fmt.Println("Migrated", filename)
	}
	return tx.Commit()
}
