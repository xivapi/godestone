package models

import (
	"time"
)

// TotalAchievementInfo represents information about a character's achievements in aggregate.
type TotalAchievementInfo struct {
	TotalAchievements      uint32
	TotalAchievementPoints uint32
}

// AchievementInfo represents information about a character's achievements.
type AchievementInfo struct {
	*TotalAchievementInfo

	Error error

	ID   uint32
	Date time.Time
}
