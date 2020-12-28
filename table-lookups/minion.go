package lookups

import (
	"strings"

	"github.com/karashiiro/godestone/pack/exports"
)

// MinionTableLookup searches the provided table for the minion that matches the provided name.
func MinionTableLookup(minionTable *exports.MinionTable, name string) *exports.Minion {
	nameLower := strings.ToLower(name)
	nameLower = strings.Replace(nameLower, "beseeltes", "beseelt", 1) // Thanks, German

	nMinions := minionTable.MinionsLength()
	for i := 0; i < nMinions; i++ {
		minion := exports.Minion{}
		minionTable.Minions(&minion, i)

		nameEn := string(minion.NameEn())
		nameDe := string(minion.NameDe())
		nameFr := string(minion.NameFr())
		nameJa := string(minion.NameJa())

		if listContains(
			nameLower,
			nameEn,
			nameDe,
			"Schwarzes "+nameDe,
			nameFr,
			nameJa,
		) {
			return &minion
		}
	}

	return nil
}
