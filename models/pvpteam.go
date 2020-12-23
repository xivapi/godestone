package models

// PVPTeamSearchResult represents basic PVP team information returned from a search.
type PVPTeamSearchResult struct {
	Error error

	Name        string
	ID          string
	DC          string
	CrestLayers *CrestLayers
}
