package lookups

import (
	"strings"

	"github.com/karashiiro/godestone/pack/exports"
)

// ItemTableLookup searches the provided table for the item that matches the provided name.
func ItemTableLookup(itemTable *exports.ItemTable, name string) *exports.Item {
	nameLower := strings.ToLower(name)

	nItems := itemTable.ItemsLength()
	for i := 0; i < nItems; i++ {
		item := exports.Item{}
		itemTable.Items(&item, i)

		nameEn := string(item.NameEn())
		nameDe := string(item.NameDe())
		nameFr := string(item.NameFr())
		nameJa := string(item.NameJa())

		nameEnLower := strings.ToLower(nameEn)
		nameDeLower := strings.ToLower(nameDe)
		nameFrLower := strings.ToLower(nameFr)
		nameJaLower := strings.ToLower(nameJa)

		if nameEnLower == nameLower || nameDeLower == nameLower || nameFrLower == nameLower || nameJaLower == nameLower {
			return &item
		}
	}

	return nil
}
