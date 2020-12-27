package lookups

import (
	"strings"

	"github.com/karashiiro/godestone/pack/exports"
)

// TribeTableLookup searches the provided table for the tribe that matches the provided name.
func TribeTableLookup(tribeTable *exports.TribeTable, name string) *exports.Tribe {
	nameLower := strings.ToLower(name)

	nTribes := tribeTable.TribesLength()
	for i := 0; i < nTribes; i++ {
		tribe := exports.Tribe{}
		tribeTable.Tribes(&tribe, i)

		nameMasculineEn := string(tribe.NameMasculineEn())
		nameMasculineDe := string(tribe.NameMasculineDe())
		nameMasculineFr := string(tribe.NameMasculineFr())
		nameMasculineJa := string(tribe.NameMasculineJa())
		nameFeminineEn := string(tribe.NameFeminineEn())
		nameFeminineDe := string(tribe.NameFeminineDe())
		nameFeminineFr := string(tribe.NameFeminineFr())
		nameFeminineJa := string(tribe.NameFeminineJa())

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
			return &tribe
		}
	}

	return nil
}
