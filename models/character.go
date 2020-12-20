package models

import "time"

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
	GearSet           *Gearset
	Gender            uint8
	GrandCompany      *GrandCompany
	GuardianDeity     uint8
	ID                uint32
	Lang              string
	Name              string
	Nameday           string
	ParseDate         *time.Time
	Portrait          string
	PvPTeamID         string
	Race              uint8
	Server            string
	Title             uint32
	TitleTop          bool
	Town              uint8
	Tribe             uint8
}

// CharacterExtended represents enriched information available about a character on The Lodestone.
type CharacterExtended struct {
	//
}
