package provider

import "github.com/xivapi/godestone/v2/provider/models"

// DataProvider represents a service implementation responsible for
// converting our scraped data into a useful form.
type DataProvider interface {
	Achievement(name string) (*models.NamedEntity, error)
	ClassJob(name string) (*models.NamedEntity, error)
	Deity(name string) (*models.NamedEntity, error)
	GrandCompany(name string) (*models.NamedEntity, error)
	Item(name string) (*models.NamedEntity, error)
	Minion(name string) (*models.NamedEntity, error)
	Mount(name string) (*models.NamedEntity, error)
	Race(name string) (*models.GenderedEntity, error)
	Reputation(name string) (*models.NamedEntity, error)
	Title(name string) (*models.TitleInternal, error)
	Town(name string) (*models.NamedEntity, error)
	Tribe(name string) (*models.GenderedEntity, error)
}
