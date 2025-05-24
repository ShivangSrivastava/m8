package core

import "time"

// Migration represents a single database schema change, either to be applied (up) or reverted (down).
// It includes metadata like version (for ordering), name (for readability),
// and optional timestamp to indicate when it was applied.
type Migration struct {
	Version   string     `json:"version"`
	AppliedAt *time.Time `json:"applied_at"`
	UpSQL     string     `json:"up_sql"`
	DownSQL   string     `json:"down_sql"`
	Name      string     `json:"name"`
}

// MigrationRepo defines an interface for interacting with the storage layer that tracks applied migrations.
// Abstracting this behavior behind an interface decouples migration logic from the underlying data store,
// making it easier to test, extend, or swap out persistence mechanisms (e.g., PostgreSQL, file, memory).
type MigrationRepo interface {
	GetAppliedMigrations() ([]Migration, error)
	ApplyMigration(Migration) error
	RevertMigration() error
}
