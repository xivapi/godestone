package provider

import "github.com/xivapi/godestone/v2/provider/models"

// DataProvider represents a service implementation responsible for
// converting our scraped data into a useful form.
type DataProvider interface {
	// Achievement returns expanded game data for the achievement with the specified name.
	Achievement(name string) (*models.NamedEntity, error)

	// ClassJob returns expanded game data for the ClassJob with the specified name.
	ClassJob(name string) (*models.NamedEntity, error)

	// Deity returns expanded game data for the deity with the specified name.
	Deity(name string) (*models.NamedEntity, error)

	// GrandCompany returns expanded game data for the Grand Company with the specified name.
	GrandCompany(name string) (*models.NamedEntity, error)

	// Item returns expanded game data for the item with the specified name.
	Item(name string) (*models.NamedEntity, error)

	// Minion returns expanded game data for the minion with the specified name.
	Minion(name string) (*models.NamedEntity, error)

	// Mount returns expanded game data for the mount with the specified name.
	Mount(name string) (*models.NamedEntity, error)

	// Race returns expanded game data for the race with the specified name.
	Race(name string) (*models.GenderedEntity, error)

	// Reputation returns expanded game data for the reputation with the specified name.
	Reputation(name string) (*models.NamedEntity, error)

	// Title returns expanded game data for the title with the specified name.
	Title(name string) (*models.TitleInternal, error)

	// Town returns expanded game data for the town with the specified name.
	Town(name string) (*models.NamedEntity, error)

	// Tribe returns expanded game data for the tribe with the specified name.
	Tribe(name string) (*models.GenderedEntity, error)
}
