package app

import (
	"github.com/ShivangSrivastava/m8/internal/core"
)

type MigrationLoader interface {
	LoadMigrations() ([]core.Migration, error)
}

type ApplyService struct {
	Repo        core.MigrationRepo
	Loader      MigrationLoader
	AppliedName []string
}

// Apply processes all available "up" migrations, filters out those already applied,
// and applies the remaining ones in order using the provided MigrationRepo.
func (a *ApplyService) Apply() error {
	all, err := a.Loader.LoadMigrations()
	if err != nil {
		return err
	}

	applied, err := a.Repo.GetAppliedMigrations()
	if err != nil {
		return err
	}
	appliedMap := map[string]bool{}

	for _, m := range applied {
		appliedMap[m.Version] = true
	}

	for _, m := range all {
		if appliedMap[m.Version] {
			continue
		}
		a.AppliedName = append(a.AppliedName, m.Name)
		if err := a.Repo.ApplyMigration(m); err != nil {
			return err
		}
	}
	return nil
}
