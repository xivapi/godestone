package provider

import "github.com/xivapi/godestone/v2/provider/models"

// DataProvider represents a service implementation responsible for
// converting our scraped data into a useful form.
type DataProvider interface {
	Achievement(name string) *models.NamedEntity
	ClassJob(name string) *models.NamedEntity
	Deity(name string) *models.NamedEntity
	GrandCompany(name string) *models.NamedEntity
	Item(name string) *models.NamedEntity
	Minion(name string) *models.NamedEntity
	Mount(name string) *models.NamedEntity
	Race(name string) *models.GenderedEntity
	Reputation(name string) *models.NamedEntity
	Title(name string) *models.TitleInternal
	Town(name string) *models.NamedEntity
	Tribe(name string) *models.GenderedEntity
}
