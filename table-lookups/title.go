package lookups

import (
	"strings"

	"github.com/karashiiro/godestone/pack/exports"
)

// TitleTableLookup searches the provided table for the title that matches the provided name.
func TitleTableLookup(titleTable *exports.TitleTable, name string) *exports.Title {
	nameLower := strings.ToLower(name)

	nTitles := titleTable.TitlesLength()
	for i := 0; i < nTitles; i++ {
		title := exports.Title{}
		titleTable.Titles(&title, i)

		nameMasculineEn := string(title.NameMasculineEn())
		nameMasculineDe := string(title.NameMasculineDe())
		nameMasculineFr := string(title.NameMasculineFr())
		nameMasculineJa := string(title.NameMasculineJa())
		nameFeminineEn := string(title.NameFeminineEn())
		nameFeminineDe := string(title.NameFeminineDe())
		nameFeminineFr := string(title.NameFeminineFr())
		nameFeminineJa := string(title.NameFeminineJa())

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
			return &title
		}
	}

	return nil
}
