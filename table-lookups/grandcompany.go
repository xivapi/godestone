package lookups

import (
	"strings"

	"github.com/karashiiro/godestone/pack/exports"
)

// GrandCompanyTableLookup searches the provided table for the Grand Company that matches the provided name.
func GrandCompanyTableLookup(grandCompanyTable *exports.GrandCompanyTable, name string) *exports.GrandCompany {
	nameLower := strings.ToLower(name)

	nGCs := grandCompanyTable.GrandCompaniesLength()
	for i := 0; i < nGCs; i++ {
		gc := exports.GrandCompany{}
		grandCompanyTable.GrandCompanies(&gc, i)

		nameEn := string(gc.NameEn())
		nameJa := string(gc.NameJa())
		nameDe := string(gc.NameDe())
		nameFr := string(gc.NameFr())

		if listContains(
			nameLower,
			nameEn,
			nameDe,
			nameFr,
			nameJa,
		) {
			return &gc
		}
	}

	return nil
}
