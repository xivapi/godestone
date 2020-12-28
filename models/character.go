package models

import (
	"time"

	"github.com/karashiiro/godestone/data/gcrank"
	"github.com/karashiiro/godestone/data/gender"
)

// Title represents a character title.
type Title struct {
	*GenderedEntity

	Prefix bool
}

// Character represents the information available about a character on The Lodestone.
type Character struct {
	ActiveClassJob    *ClassJob
	Avatar            string
	Bio               string
	ClassJobs         []*ClassJob
	ClassJobBozjan    *ClassJobBozja
	ClassJobElemental *ClassJobEureka
	DC                string
	FreeCompanyID     string
	FreeCompanyName   string
	GearSet           *GearSet
	Gender            gender.Gender
	GrandCompanyInfo  *GrandCompanyInfo
	GuardianDeity     *NamedEntity
	ID                uint32
	Name              string
	Nameday           string
	ParseDate         time.Time
	Portrait          string
	PvPTeamID         string
	Race              *GenderedEntity
	Title             *Title
	Town              *NamedEntity
	Tribe             *GenderedEntity
	World             string
}

// CharacterSearchResult contains data from the character search page about a character.
type CharacterSearchResult struct {
	Error error

	Avatar   string
	ID       uint32
	Lang     string
	Name     string
	Rank     gcrank.GCRank
	RankIcon string
	World    string
	DC       string
}
