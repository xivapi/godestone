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

		nameMasculineEn := string(race.NameMasculineEn())
		nameMasculineDe := string(race.NameMasculineDe())
		nameMasculineFr := string(race.NameMasculineFr())
		nameMasculineJa := string(race.NameMasculineJa())
		nameFeminineEn := string(race.NameFeminineEn())
		nameFeminineDe := string(race.NameFeminineDe())
		nameFeminineFr := string(race.NameFeminineFr())
		nameFeminineJa := string(race.NameFeminineJa())

		nameMasculineEnLower := strings.ToLower(nameMasculineEn)
		nameMasculineDeLower := strings.ToLower(nameMasculineDe)
		nameMasculineFrLower := strings.ToLower(nameMasculineFr)
		nameMasculineJaLower := strings.ToLower(nameMasculineJa)
		nameFeminineEnLower := strings.ToLower(nameFeminineEn)
		nameFeminineDeLower := strings.ToLower(nameFeminineDe)
		nameFeminineFrLower := strings.ToLower(nameFeminineFr)
		nameFeminineJaLower := strings.ToLower(nameFeminineJa)

		if nameMasculineEnLower == nameLower ||
			nameMasculineDeLower == nameLower ||
			nameMasculineFrLower == nameLower ||
			nameMasculineJaLower == nameLower ||
			nameFeminineEnLower == nameLower ||
			nameFeminineDeLower == nameLower ||
			nameFeminineFrLower == nameLower ||
			nameFeminineJaLower == nameLower {
			return &race
		}
	}

	return nil
}
