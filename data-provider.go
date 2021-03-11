package godestone

import (
	"github.com/karashiiro/godestone/pack/exports"
)

// DataProvider represents a service implementation responsible for
// converting our scraped data into a useful form.
type DataProvider interface {
	ClassJob(name string) *ClassJob
	Deity(name string) *NamedEntity
	GrandCompany(name string) *NamedEntity
	Item(name string) *exports.Item
	Minion(name string) *Minion
	Mount(name string) *Mount
	Race(name string) *GenderedEntity
	Reputation(name string) *NamedEntity
	Title(name string) *Title
	Town(name string) *NamedEntity
	Tribe(name string) *GenderedEntity
}
