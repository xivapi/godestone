package models

import (
	"time"
)

// AchievementInfo represents information about a single achievement.
type AchievementInfo struct {
	Error error

	ID   uint32
	Date time.Time
}
