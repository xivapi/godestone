package models

import (
	"github.com/karashiiro/godestone/data/gcrank"
)

// GrandCompanyInfo represents Grand Company information about a character.
type GrandCompanyInfo struct {
	GrandCompany *NamedEntity
	RankID       gcrank.GCRank
}
