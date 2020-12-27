package lookups

import (
	"strings"

	"github.com/karashiiro/godestone/pack/exports"
)

// TownTableLookup searches the provided table for the town that matches the provided name.
func TownTableLookup(townTable *exports.TownTable, name string) *exports.Town {
	nameLower := strings.ToLower(name)

	nTowns := townTable.TownsLength()
	for i := 0; i < nTowns; i++ {
		town := exports.Town{}
		townTable.Towns(&town, i)

		nameEn := string(town.NameEn())
		nameDe := string(town.NameDe())
		nameFr := string(town.NameFr())
		nameJa := string(town.NameJa())

		nameEnLower := strings.ToLower(nameEn)
		nameDeLower := strings.ToLower(nameDe)
		nameFrLower := strings.ToLower(nameFr)
		nameJaLower := strings.ToLower(nameJa)

		if nameEnLower == nameLower || nameDeLower == nameLower || nameFrLower == nameLower || nameJaLower == nameLower {
			return &town
		}
	}

	return nil
}
