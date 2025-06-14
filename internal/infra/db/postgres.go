package db

import (
	"database/sql"
	"errors"
	"time"

	"github.com/ShivangSrivastava/m8/internal/core"
)

type DBRepo struct {
	DB *sql.DB
}

// ensureTable creates the schema_migrations table if it doesn't exist.
// This guarantees that the migration tracking infrastructure is present before any operations.
// Doing this lazily on-demand avoids requiring manual setup and simplifies usage.
func (d *DBRepo) ensureTable() error {
	_, err := d.DB.Exec(`
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
func (d *DBRepo) GetAppliedMigrations() ([]core.Migration, error) {
	if err := d.ensureTable(); err != nil {
		return nil, err
	}

	rows, err := d.DB.Query("SELECT version, applied_at FROM schema_migrations ORDER BY applied_at")
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

// GetLatestMigration fetches the most recently applied migration.
// Used for determining which migration to revert in `down`.
// Returns error if no migrations have been applied yet.
func (d *DBRepo) GetLatestMigration() (core.Migration, error) {
	if err := d.ensureTable(); err != nil {
		return core.Migration{}, err
	}

	rows, err := d.DB.Query("SELECT version, applied_at FROM schema_migrations ORDER BY applied_at DESC LIMIT 1")
	if err != nil {
		return core.Migration{}, err
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
	if len(result) == 0 {
		return core.Migration{}, errors.New("no migration is to revert")
	}
	return result[0], nil
}

// ApplyMigration runs the migration’s UpSQL inside a transaction to ensure atomicity.
// It records the migration version in schema_migrations upon success,
// so future runs can detect it as applied and avoid reapplying.
// Using transactions protects the database from partial or corrupted states in case of failure.
func (d *DBRepo) ApplyMigration(m core.Migration) error {
	tx, err := d.DB.Begin()
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

// RevertMigration rolls back the latest migration using DownSQL.
// Removes its record from schema_migrations to mark it as unapplied.
// Ensures database state consistency with transaction safety.
func (d *DBRepo) RevertMigration(m core.Migration) error {
	tx, err := d.DB.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec(m.DownSQL)
	if err != nil {
		return err
	}

	_, err = tx.Exec("DELETE FROM schema_migrations WHERE version=$1", m.Version)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
