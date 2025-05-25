package app

type Status string

// Defines migration states: "applied" and "pending".
const (
	Pending Status = "pending"
	Applied Status = "applied"
)

// `MigrationStatus` pairs a migration's name with its current status.
type MigrationStatus struct {
	Status Status
	Name   string
}

// `GetStatus` compares loaded vs applied migrations to return a full status list.
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
