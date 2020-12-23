package models

import (
	"time"
)

// AchievementInfo represents information about a character's achievements.
type AchievementInfo struct {
	Error error

	TotalAchievements      uint32
	TotalAchievementPoints uint32

	ID   uint32
	Date time.Time
}
