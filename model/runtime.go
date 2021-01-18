package model

import "time"

// RuntimeUpgrade used to create JSON file for temp file
type RuntimeUpgrade struct {
	StartTime     time.Time
	EndTime       time.Time
	BusinessID    int
	UpgradeObject string
}
