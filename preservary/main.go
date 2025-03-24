package main

import (
	"log"
	"preservary/config"
	// "github.com/go-co-op/gocron"
)

func main() {
	// Load the configuration file
	config, err := config.LoadConfig("/config.json") // Replace with the path to your JSON file

	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Get the locations from the config
	locations, err := config.GetLocations()

	if err != nil {
		log.Fatalf("Error getting locations: %v", err)
	}

	// Print the locations
	for i, loc := range locations {
		log.Printf("Location %d: %+v\n", i+1, loc)
	}
}
