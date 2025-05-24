package core

import "time"

type Migration struct {
	Version   string     `json:"version"`
	AppliedAt *time.Time `json:"applied_at"`
	UpSQL     string     `json:"up_sql"`
	DownSQL   string     `json:"down_sql"`
	Name      string     `json:"name"`
}

type MigrationRepo interface {
	GetAppliedMigrations() ([]Migration, error)
	ApplyMigration(Migration) error
	RevertMigration() error
}
