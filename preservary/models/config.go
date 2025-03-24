package models

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	// LogLevel     string         `json:"LogLevel"`
	// BackupPolicy string         `json:"BackupPolicy"`
	Locations []LocationJSON `json:"Locations"`
}

func LoadConfig(filePath string) (*Config, error) {
	file, err := os.Open(filePath)

	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}

	defer file.Close()

	var config Config

	if err := json.NewDecoder(file).Decode(&config); err != nil {
		return nil, fmt.Errorf("error decoding JSON: %w", err)
	}

	return &config, nil
}

func (c *Config) GetLocations() ([]Location, error) {
	var locations []Location

	for _, locJSON := range c.Locations {
		location, err := NewLocation(
			locJSON.SourcePath,
			locJSON.DestinationPath,
			locJSON.ComposePath,
			locJSON.BackupInterval,
			locJSON.Archive,
		)

		if err != nil {
			return nil, fmt.Errorf("error creating Location: %w", err)
		}

		locations = append(locations, *location)
	}

	return locations, nil
}
