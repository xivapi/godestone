package models

// CWLSSearchResult represents basic CWLS information returned from a search.
type CWLSSearchResult struct {
	Error error

	Name          string
	ID            string
	DC            string
	ActiveMembers uint32
}
