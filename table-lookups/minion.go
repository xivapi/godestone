package lookups

import (
	"strings"

	"github.com/karashiiro/godestone/pack/exports"
)

// MinionTableLookup searches the provided table for the minion that matches the provided name.
func MinionTableLookup(minionTable *exports.MinionTable, name string, lang string) *exports.Minion {
	nameLower := strings.ToLower(name)

	// Thanks, German
	// If anyone knows how to properly handle this, a PR would be more than welcome.
	if lang == "de" {
		nameLower = strings.Replace(nameLower, "blaublütiger ", "baby-", 1)
		nameLower = strings.Replace(nameLower, "es ", " ", 1)
		if !strings.HasPrefix(nameLower, "seite") {
			nameLower = strings.Replace(nameLower, "e ", " ", 1)
		}
		nameLower = strings.Replace(nameLower, "er ", " ", 1)
		nameLower = strings.Replace(nameLower, " d ", " der ", 1)
		if strings.Contains(nameLower, "chocobo-küken") {
			parts := strings.Split(nameLower, " ")
			nameLower = parts[len(parts)-1]
		}
	}

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
			nameFr,
			nameJa,
		) {
			return &minion
		}
	}

	return nil
}
