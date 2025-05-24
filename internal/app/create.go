package app

import (
	"regexp"
	"strings"
	"time"
)

// MigrationCreator defines the interface for generating migration files.
// It allows decoupling of migration file logic from the CreateService, making it testable and extensible.
type MigrationCreator interface {
	CreateMigrations(version, name string) error
}

// CreateService encapsulates logic for generating timestamped, sanitized migration files.
// It delegates file creation to a MigrationCreator and handles name formatting and versioning.
type CreateService struct {
	Name    string
	Creator MigrationCreator
}

// getTS returns the current timestamp in a format suitable for versioning migration files.
// This ensures consistent, sortable naming across all generated migrations.
func (c *CreateService) getTS() string {
	return time.Now().Format("20060102150405")
}

// sanitizeName removes unsupported characters from the migration name,
// ensuring it produces a filesystem-safe identifier composed of alphanumeric and underscore segments.
func (c *CreateService) sanitizeName() (string, error) {
	r, err := regexp.Compile("[a-zA-Z0-9_]+")
	if err != nil {
		return "", err
	}
	s := r.FindAllString(c.Name, -1)
	return strings.Join(s, "_"), nil
}

// CreateFile generates a timestamped and sanitized migration file pair via the configured Creator.
// It handles naming consistency and version generation, reducing manual migration errors.
func (c *CreateService) CreateFile() error {
	version := c.getTS()
	sanitizedName, err := c.sanitizeName()
	if err != nil {
		return err
	}
	return c.Creator.CreateMigrations(version, sanitizedName)
}
