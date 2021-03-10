package models

import (
	"time"

	"github.com/karashiiro/godestone/data/gcrank"
)

// CWLS represents basic CWLS information.
type CWLS struct {
	Name      string
	DC        string
	ID        string
	ParseDate time.Time
	Members   []*CWLSMember
}

// CWLSMember represents information about a CWLS member.
type CWLSMember struct {
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

// CWLSSearchResult represents basic CWLS information returned from a search.
type CWLSSearchResult struct {
	Error error `json:"-"`

	Name          string
	ID            string
	DC            string
	ActiveMembers uint32
}
