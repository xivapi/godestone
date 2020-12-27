package lookups

import (
	"strings"

	"github.com/karashiiro/godestone/pack/exports"
)

// RaceTableLookup searches the provided table for the race that matches the provided name.
func RaceTableLookup(raceTable *exports.RaceTable, name string) *exports.Race {
	nameLower := strings.ToLower(name)

	nRaces := raceTable.RacesLength()
	for i := 0; i < nRaces; i++ {
		race := exports.Race{}
		raceTable.Races(&race, i)

		nameEn := string(race.NameEn())
		nameDe := string(race.NameDe())
		nameFr := string(race.NameFr())
		nameJa := string(race.NameJa())

		nameEnLower := strings.ToLower(nameEn)
		nameDeLower := strings.ToLower(nameDe)
		nameFrLower := strings.ToLower(nameFr)
		nameJaLower := strings.ToLower(nameJa)

		if nameEnLower == nameLower || nameDeLower == nameLower || nameFrLower == nameLower || nameJaLower == nameLower {
			return &race
		}
	}

	return nil
}
