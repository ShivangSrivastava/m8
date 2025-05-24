package db

import (
	"database/sql"
	"time"

	"github.com/ShivangSrivastava/m8/internal/core"
)

type DBRepo struct {
	DB *sql.DB
}

func (p *DBRepo) ensureTable() error {
	_, err := p.DB.Exec(`
		CREATE TABLE IF NOT EXISTS schema_migrations (
		version TEXT PRIMARY KEY,
		applied_at TIMESTAMPTZ DEFAULT NOW()
		)`)
	return err
}

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
