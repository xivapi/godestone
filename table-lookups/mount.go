package lookups

import (
	"strings"

	"github.com/karashiiro/godestone/pack/exports"
)

// MountTableLookup searches the provided table for the mount that matches the provided name.
func MountTableLookup(mountTable *exports.MountTable, name string) *exports.Mount {
	nameLower := strings.ToLower(name)

	nMounts := mountTable.MountsLength()
	for i := 0; i < nMounts; i++ {
		mount := exports.Mount{}
		mountTable.Mounts(&mount, i)

		nameEn := string(mount.NameEn())
		nameDe := string(mount.NameDe())
		nameFr := string(mount.NameFr())
		nameJa := string(mount.NameJa())

		if listContains(
			nameLower,
			nameEn,
			nameDe,
			nameFr,
			nameJa,
		) {
			return &mount
		}
	}

	return nil
}
