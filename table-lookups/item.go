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

		if listContains(
			nameLower,
			nameEn,
			nameDe,
			nameFr,
			nameJa,
		) {
			return &item
		}
	}

	return nil
}
