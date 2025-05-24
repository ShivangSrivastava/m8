package app

import (
	"regexp"
	"strings"
	"time"
)

type MigrationCreator interface {
	CreateMigrations(version, name string) error
}
type CreateService struct {
	Name    string
	Creator MigrationCreator
}

func (c *CreateService) getTS() string {
	return time.Now().Format("20060102150405")
}

func (c *CreateService) sanitizeName() (string, error) {
	r, err := regexp.Compile("[a-zA-Z0-9_]+")
	if err != nil {
		return "", err
	}
	s := r.FindAllString(c.Name, -1)
	return strings.Join(s, "_"), nil
}

func (c *CreateService) CreateFile() error {
	version := c.getTS()
	sanitizedName, err := c.sanitizeName()
	if err != nil {
		return err
	}
	return c.Creator.CreateMigrations(version, sanitizedName)
}
