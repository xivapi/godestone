package godestone

// DataProvider represents a service implementation responsible for
// converting our scraped data into a useful form.
type DataProvider interface {
	Achievement(name string) *NamedEntity
	ClassJob(name string) *NamedEntity
	Deity(name string) *NamedEntity
	GrandCompany(name string) *NamedEntity
	Item(name string) *NamedEntity
	Minion(name string) *NamedEntity
	Mount(name string) *NamedEntity
	Race(name string) *GenderedEntity
	Reputation(name string) *NamedEntity
	Title(name string) *Title
	Town(name string) *NamedEntity
	Tribe(name string) *GenderedEntity
}
