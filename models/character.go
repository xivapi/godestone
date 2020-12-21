package models

import (
	"time"

	"github.com/karashiiro/godestone/data/deity"
	"github.com/karashiiro/godestone/data/gcrank"
	"github.com/karashiiro/godestone/data/gender"
	"github.com/karashiiro/godestone/data/race"
	"github.com/karashiiro/godestone/data/town"
	"github.com/karashiiro/godestone/data/tribe"
)

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
	GrandCompany      *GrandCompanyInfo
	GuardianDeity     deity.GuardianDeity
	ID                uint32
	Name              string
	Nameday           string
	ParseDate         *time.Time
	Portrait          string
	PvPTeamID         string
	Race              race.Race
	Server            string
	Title             uint32
	TitleTop          bool
	Town              town.Town
	Tribe             tribe.Tribe
}

// CharacterExtended represents enriched information available about a character on The Lodestone.
type CharacterExtended struct {
	//
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
	Server   string
}
