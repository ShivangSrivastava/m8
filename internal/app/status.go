package app

type Status string

const (
	Pending Status = "pending"
	Applied Status = "applied"
)

type MigrationStatus struct {
	Status Status
	Name   string
}

func (a *ApplyService) GetStatus() ([]MigrationStatus, error) {
	all, err := a.Loader.LoadMigrations()
	if err != nil {
		return nil, err
	}

	applied, err := a.Repo.GetAppliedMigrations()
	if err != nil {
		return nil, err
	}
	appliedMap := map[string]bool{}

	for _, m := range applied {
		appliedMap[m.Version] = true
	}

	var result []MigrationStatus
	for _, m := range all {
		if appliedMap[m.Version] {
			result = append(result, MigrationStatus{
				Name:   m.Name,
				Status: Applied,
			})
		} else {
			result = append(result, MigrationStatus{
				Name:   m.Name,
				Status: Pending,
			})
		}
	}
	return result, nil
}
