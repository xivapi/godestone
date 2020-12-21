package models

import (
	"time"
)

// AchievementInfo represents information about a single achievement.
type AchievementInfo struct {
	Date time.Time
	ID   uint32
}

// Achievements represents information about all of a character's achievements.
type Achievements struct {
	List   []*AchievementInfo
	Points uint32
}
