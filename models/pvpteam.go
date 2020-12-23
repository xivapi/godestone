package models

import (
	"time"

	"github.com/karashiiro/godestone/data/gcrank"
)

// PVPTeam represents information about a PVP team.
type PVPTeam struct {
	Name        string
	ID          string
	DC          string
	ParseDate   time.Time
	Formed      time.Time
	CrestLayers *CrestLayers
	Members     []*PVPTeamMember
}

// PVPTeamMember represents information about a PVP team member.
type PVPTeamMember struct {
	Avatar   string
	ID       uint32
	Name     string
	Matches  uint32
	Rank     gcrank.GCRank
	RankIcon string
	World    string
	DC       string
}

// PVPTeamSearchResult represents basic PVP team information returned from a search.
type PVPTeamSearchResult struct {
	Error error

	Name        string
	ID          string
	DC          string
	CrestLayers *CrestLayers
}
