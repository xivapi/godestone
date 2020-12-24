package search

import (
	"fmt"
	"strings"
)

// PVPTeamSearchOrder represents the search result ordering of a Lodestone CWLS search.
type PVPTeamSearchOrder uint8

// Search ordering for PVP Team searches.
const (
	OrderPVPTeamNameAToZ PVPTeamSearchOrder = iota + 1
	OrderPVPTeamNameZToA
)

// PVPTeamOptions defines extra search information that can help to narrow down a PVP team search.
type PVPTeamOptions struct {
	Name                      string
	DC                        string
	Order                     PVPTeamSearchOrder
	CommunityFinderRecruiting bool
}

// BuildURI returns a constructed URI for the provided search options.
func (s *PVPTeamOptions) BuildURI(lang string) string {
	uriFormat := "https://%s.finalfantasyxiv.com/lodestone/pvpteam/?q=%s&dcname=%s&cf_public=%d&order=%d"

	name := strings.Replace(s.Name, " ", "%20", -1)

	cfPublic := 0
	if s.CommunityFinderRecruiting {
		cfPublic = 1
	}

	builtURI := fmt.Sprintf(uriFormat, lang, name, s.DC, cfPublic, s.Order)
	return builtURI
}
