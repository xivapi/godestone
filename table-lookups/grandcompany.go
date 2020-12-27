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

		nameEnLower := strings.ToLower(nameEn)
		nameJaLower := strings.ToLower(nameJa)
		nameDeLower := strings.ToLower(nameDe)
		nameFrLower := strings.ToLower(nameFr)

		// This is different for GCs compared to the other tables because of inconsistency in the presence/absence
		// of definite articles in various languages.
		if strings.Contains(nameEnLower, nameLower) || strings.Contains(nameJaLower, nameLower) || strings.Contains(nameDeLower, nameLower) || strings.Contains(nameFrLower, nameLower) {
			return &gc
		}
	}

	return nil
}
