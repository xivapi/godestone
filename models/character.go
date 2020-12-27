package models

import (
	"time"

	"github.com/karashiiro/godestone/data/gcrank"
	"github.com/karashiiro/godestone/data/gender"
	"github.com/karashiiro/godestone/data/town"
)

// Title represents a title that a character can have.
type Title struct {
	ID     uint32
	Name   string
	Prefix bool

	NameMasculineEN string
	NameMasculineJA string
	NameMasculineDE string
	NameMasculineFR string
	NameFeminineEN  string
	NameFeminineJA  string
	NameFeminineDE  string
	NameFeminineFR  string
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
	Race              *NamedEntity
	Title             *Title
	Town              *struct {
		Name town.Town
		Icon string
	}
	Tribe *NamedEntity
	World string
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
