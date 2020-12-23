package models

// LinkshellSearchResult represents basic linkshell information returned from a search.
type LinkshellSearchResult struct {
	Error error

	Name          string
	ID            string
	World         string
	DC            string
	ActiveMembers uint32
}
