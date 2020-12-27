package lookups

import (
	"strings"

	"github.com/karashiiro/godestone/pack/exports"
)

// DeityTableLookup searches the provided table for the deity that matches the provided name.
func DeityTableLookup(deityTable *exports.DeityTable, name string) *exports.Deity {
	nameLower := strings.ToLower(name)

	nDeities := deityTable.DeitiesLength()
	for i := 0; i < nDeities; i++ {
		deity := exports.Deity{}
		deityTable.Deities(&deity, i)

		nameEn := string(deity.NameEn())
		nameDe := string(deity.NameDe())
		nameFr := string(deity.NameFr())
		nameJa := string(deity.NameJa())

		if listContains(
			nameLower,
			nameEn,
			nameDe,
			nameFr,
			nameJa,
		) {
			return &deity
		}
	}

	return nil
}
