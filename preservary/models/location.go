package models

import (
	"fmt"
	"strconv"
	"strings"
)

type Location struct {
	SourcePath              string
	DestinationPath         string
	ComposePath             string
	BackupIntervalInMinutes int
	Archive                 bool
}

type LocationJSON struct {
	SourcePath      string `json:"SourcePath"`
	DestinationPath string `json:"DestinationPath"`
	ComposePath     string `json:"ComposePath"`
	BackupInterval  string `json:"BackupInterval"`
	Archive         bool   `json:"Archive"`
}

func NewLocation(sourcePath, destinationPath, composePath, backupInterval string, archive bool) (*Location, error) {
	minutes, err := parseBackupInterval(backupInterval)

	if err != nil {
		return nil, err
	}

	return &Location{
		SourcePath:              sourcePath,
		DestinationPath:         destinationPath,
		ComposePath:             composePath,
		BackupIntervalInMinutes: minutes,
		Archive:                 archive,
	}, nil
}

func parseBackupInterval(interval string) (int, error) {
	const (
		MinutesPerHour = 60
		MinutesPerDay  = 24 * MinutesPerHour
	)

	if interval == "" {
		return 0, fmt.Errorf("BackupInterval cannot be empty")
	}

	if strings.HasSuffix(interval, "d") {
		days, err := strconv.Atoi(strings.TrimSuffix(interval, "d"))

		if err != nil {
			return 0, fmt.Errorf("invalid day format: %s", interval)
		}

		return days * MinutesPerDay, nil
	}

	if strings.HasSuffix(interval, "h") {
		hours, err := strconv.Atoi(strings.TrimSuffix(interval, "h"))

		if err != nil {
			return 0, fmt.Errorf("invalid hour format: %s", interval)
		}

		return hours * MinutesPerHour, nil
	}

	if strings.HasSuffix(interval, "m") {
		minutes, err := strconv.Atoi(strings.TrimSuffix(interval, "m"))

		if err != nil {
			return 0, fmt.Errorf("invalid minute format: %s", interval)
		}

		return minutes, nil
	}

	return 0, fmt.Errorf("invalid BackupInterval format: %s", interval)
}
