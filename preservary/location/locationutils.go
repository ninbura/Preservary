package location

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

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

func validatePath(path string, childCheck bool) (bool, error) {
	info, err := os.Stat(path)

	if os.IsNotExist(err) {
		return false, fmt.Errorf("path '%s' does not exist", path)
	}

	if err != nil {
		return false, fmt.Errorf("error accessing path '%s': %v", path, err)
	}

	if !info.IsDir() {
		return false, fmt.Errorf("path '%s' is not a directory", path)
	}

	if !childCheck {
		return true, nil
	}

	dir, err := os.Open(path)

	if err != nil {
		return false, fmt.Errorf("error opening directory '%s': %v", path, err)
	}

	defer dir.Close()

	children, err := dir.Readdirnames(1)

	if err != nil {
		return false, fmt.Errorf("error reading directory '%s': %v", path, err)
	}

	if len(children) > 0 {
		return true, nil
	}

	return false, fmt.Errorf("path '%s' exists but has no children", path)
}

func processLocation(loc *Location) {
	log.Printf("nice")
}

// func ScheduleLocations
