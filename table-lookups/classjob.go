package lookups

import (
	"strings"

	"github.com/karashiiro/godestone/pack/exports"
)

// ClassJobTableLookup searches the provided table for the deity that matches the provided name.
func ClassJobTableLookup(classJobTable *exports.ClassJobTable, name string) *exports.ClassJob {
	nameLower := strings.ToLower(name)

	nClassJobs := classJobTable.ClassJobsLength()
	for i := 0; i < nClassJobs; i++ {
		cj := exports.ClassJob{}
		classJobTable.ClassJobs(&cj, i)

		nameEn := string(cj.NameEn())
		nameDe := string(cj.NameDe())
		nameFr := string(cj.NameFr())
		nameJa := string(cj.NameJa())

		if listContains(
			nameLower,
			nameEn,
			nameDe,
			nameFr,
			nameJa,
		) {
			return &cj
		}
	}

	return nil
}
