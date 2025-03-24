package config

import (
	"encoding/json"
	"fmt"
	"os"
	"preservary/location"
)

type Config struct {
	// LogLevel     string         `json:"LogLevel"`
	// BackupPolicy string         `json:"BackupPolicy"`
	Locations []location.LocationJSON `json:"Locations"`
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

func (c *Config) GetLocations() ([]location.Location, error) {
	var locs []location.Location

	for _, locJSON := range c.Locations {
		location, err := locJSON.NewLocation()

		if err != nil {
			return nil, fmt.Errorf("error creating Location: %w", err)
		}

		locs = append(locs, *location)
	}

	return locs, nil
}
