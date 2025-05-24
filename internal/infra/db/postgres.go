package db

import (
	"database/sql"
	"time"

	"github.com/ShivangSrivastava/m8/internal/core"
)

type DBRepo struct {
	DB *sql.DB
}

// ensureTable creates the schema_migrations table if it doesn't exist.
// This guarantees that the migration tracking infrastructure is present before any operations.
// Doing this lazily on-demand avoids requiring manual setup and simplifies usage.
func (p *DBRepo) ensureTable() error {
	_, err := p.DB.Exec(`
		CREATE TABLE IF NOT EXISTS schema_migrations (
		version TEXT PRIMARY KEY,
		applied_at TIMESTAMPTZ DEFAULT NOW()
		)`)
	return err
}

// GetAppliedMigrations returns all migrations recorded as applied in the database,
// sorted by their applied timestamp to preserve the original application order.
// This information is critical for the system to know which migrations to skip,
// enabling idempotent migration application and safe rollbacks.
func (p *DBRepo) GetAppliedMigrations() ([]core.Migration, error) {
	if err := p.ensureTable(); err != nil {
		return nil, err
	}

	rows, err := p.DB.Query("SELECT version, applied_at FROM schema_migrations ORDER BY applied_at")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []core.Migration

	for rows.Next() {
		var m core.Migration
		var at time.Time
		rows.Scan(&m.Version, &at)
		m.AppliedAt = &at
		result = append(result, m)
	}

	return result, nil
}

// ApplyMigration runs the migration’s UpSQL inside a transaction to ensure atomicity.
// It records the migration version in schema_migrations upon success,
// so future runs can detect it as applied and avoid reapplying.
// Using transactions protects the database from partial or corrupted states in case of failure.
func (p *DBRepo) ApplyMigration(m core.Migration) error {
	tx, err := p.DB.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec(m.UpSQL)
	if err != nil {
		return err
	}

	_, err = tx.Exec("INSERT INTO schema_migrations (version) VALUES ($1)", m.Version)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (p *DBRepo) RevertMigration() error {
	tx, err := p.DB.Begin()
	if err != nil {
		return err
	}
	return tx.Commit()
}
