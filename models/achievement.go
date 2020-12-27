package models

import (
	"time"
)

// AllAchievementInfo represents information about a character's achievements in aggregate.
type AllAchievementInfo struct {
	Private                bool
	TotalAchievements      uint32
	TotalAchievementPoints uint32
}

// AchievementInfo represents information about a character's achievements.
type AchievementInfo struct {
	*AllAchievementInfo

	Error error

	Name string
	ID   uint32
	Date time.Time

	NameEN string
	NameJA string
	NameDE string
	NameFR string
}
