package location

import (
	"preservary/shared"
	"time"
)

type LocationJSON struct {
	SourcePath         string   `json:"SourcePath"`
	SourcePaths        []string `json:"SourcePaths"`
	DestinationPath    string   `json:"DestinationPath"`
	DestinationPaths   []string `json:"DestinationPaths"`
	ComposePath        string   `json:"ComposePath"`
	StartDay           string   `json:"StartDate"`
	StartTime          string   `json:"StartTime"`
	BackupInterval     string   `json:"BackupInterval"`
	ForceInitialBackup bool     `json:"ForceInitialBackup"`
	Archive            bool     `json:"Archive"`
}

type Location struct {
	SourcePath              string
	SourcePaths             []string
	DestinationPath         string
	DestinationPaths        []string
	ComposePath             string
	StartDay                string
	StartDateTime           time.Time
	BackupIntervalInMinutes int
	ForceInitialBackup      bool
	Archive                 bool
}

func (locJson *LocationJSON) NewLocation() (*Location, error) {
	startDay := shared.CalculateDay(locJson.StartDay)
	minutes, err := parseBackupInterval(locJson.BackupInterval)

	if err != nil {
		return nil, err
	}

	return &Location{
		SourcePath:              locJson.SourcePath,
		SourcePaths:             locJson.SourcePaths,
		DestinationPath:         locJson.DestinationPath,
		DestinationPaths:        locJson.DestinationPaths,
		ComposePath:             locJson.ComposePath,
		StartDay:                startDay,
		StartDateTime:           time.Now(),
		BackupIntervalInMinutes: minutes,
		Archive:                 locJson.Archive,
	}, nil
}
