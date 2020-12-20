package models

import (
	"github.com/karashiiro/godestone/data/gcrank"
	"github.com/karashiiro/godestone/data/grandcompany"
)

// GrandCompanyInfo represents Grand Company information about a character.
type GrandCompanyInfo struct {
	NameID grandcompany.GrandCompany
	RankID gcrank.GCRank
}
