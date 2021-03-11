package godestone

import (
	"github.com/karashiiro/godestone/models"
	"github.com/karashiiro/godestone/pack/exports"
)

// DataProvider represents a service implementation responsible for
// converting our scraped data into a useful form.
type DataProvider interface {
	ClassJob(name string) *models.ClassJob
	Deity(name string) *models.NamedEntity
	GrandCompany(name string) *models.NamedEntity
	Item(name string) *exports.Item    // Do something about this
	Minion(name string) *models.Minion // Or all of these
	Mount(name string) *models.Mount
	Race(name string) *models.GenderedEntity
	Reputation(name string) *models.NamedEntity
	Title(name string) *models.Title
	Town(name string) *models.NamedEntity
	Tribe(name string) *models.GenderedEntity
}
