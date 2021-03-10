package models

import (
	"time"

	"github.com/karashiiro/godestone/data/gcrank"
)

// Linkshell represents basic linkshell information.
type Linkshell struct {
	Name      string
	ID        string
	ParseDate time.Time
	Members   []*LinkshellMember
}

// LinkshellMember represents information about a linkshell member.
type LinkshellMember struct {
	Avatar            string
	ID                uint32
	Name              string
	LinkshellRank     string
	LinkshellRankIcon string
	Rank              gcrank.GCRank
	RankIcon          string
	World             string
	DC                string
}

// LinkshellSearchResult represents basic linkshell information returned from a search.
type LinkshellSearchResult struct {
	Error error `json:"-"`

	Name          string
	ID            string
	World         string
	DC            string
	ActiveMembers uint32
}
