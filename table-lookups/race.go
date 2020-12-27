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

		if listContains(
			nameLower,
			nameMasculineEn,
			nameMasculineDe,
			nameMasculineFr,
			nameMasculineJa,
			nameFeminineEn,
			nameFeminineDe,
			nameFeminineFr,
			nameFeminineJa,
		) {
			return &race
		}
	}

	return nil
}
