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

		nameEn := string(tribe.NameEn())
		nameDe := string(tribe.NameDe())
		nameFr := string(tribe.NameFr())
		nameJa := string(tribe.NameJa())

		nameEnLower := strings.ToLower(nameEn)
		nameDeLower := strings.ToLower(nameDe)
		nameFrLower := strings.ToLower(nameFr)
		nameJaLower := strings.ToLower(nameJa)

		if nameEnLower == nameLower || nameDeLower == nameLower || nameFrLower == nameLower || nameJaLower == nameLower {
			return &tribe
		}
	}

	return nil
}
