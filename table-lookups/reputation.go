package lookups

import (
	"strings"

	"github.com/karashiiro/godestone/pack/exports"
)

// ReputationTableLookup searches the provided table for the reputation that matches the provided name.
func ReputationTableLookup(reputationTable *exports.ReputationTable, name string) *exports.Reputation {
	nameLower := strings.ToLower(name)

	nReputations := reputationTable.ReputationsLength()
	for i := 0; i < nReputations; i++ {
		reputation := exports.Reputation{}
		reputationTable.Reputations(&reputation, i)

		nameEn := string(reputation.NameEn())
		nameDe := string(reputation.NameDe())
		nameFr := string(reputation.NameFr())
		nameJa := string(reputation.NameJa())

		if listContains(
			nameLower,
			nameEn,
			nameDe,
			nameFr,
			nameJa,
		) {
			return &reputation
		}
	}

	return nil
}
