package search

import (
	"fmt"
	"strings"
)

// LinkshellSearchOrder represents the search result ordering of a Lodestone CWLS search.
type LinkshellSearchOrder uint8

// Search ordering for linkshell and CWLS searches.
const (
	OrderLinkshellNameAToZ LinkshellSearchOrder = iota + 1
	OrderLinkshellNameZToA
	OrderLinkshellMembershipHighToLow
	OrderLinkshellMembershipLowToHigh
)

// LinkshellOptions defines extra search information that can help to narrow down a linkshell search.
type LinkshellOptions struct {
	Name                      string
	World                     string
	DC                        string
	Order                     LinkshellSearchOrder
	ActiveMembers             ActiveMemberRange
	CommunityFinderRecruiting bool
}

// BuildURI returns a constructed URI for the provided search options.
func (s *LinkshellOptions) BuildURI() string {
	uriFormat := "https://na.finalfantasyxiv.com/lodestone/linkshell/?q=%s&worldname=%s&character_count=%s&cf_public=%d&order=%d"

	name := strings.Replace(s.Name, " ", "%20", -1)

	worldDC := parseWorldDC(s.World, s.DC)

	cfPublic := 0
	if s.CommunityFinderRecruiting {
		cfPublic = 1
	}

	builtURI := fmt.Sprintf(uriFormat, name, worldDC, s.ActiveMembers, cfPublic, s.Order)
	return builtURI
}

// CWLSOptions defines extra search information that can help to narrow down a CWLS search.
type CWLSOptions struct {
	Name                      string
	DC                        string
	Order                     LinkshellSearchOrder
	ActiveMembers             ActiveMemberRange
	CommunityFinderRecruiting bool
}

// BuildURI returns a constructed URI for the provided search options.
func (s *CWLSOptions) BuildURI() string {
	uriFormat := "https://na.finalfantasyxiv.com/lodestone/crossworld_linkshell/?q=%s&dcname=%s&character_count=%s&cf_public=%d&order=%d"

	name := strings.Replace(s.Name, " ", "%20", -1)

	cfPublic := 0
	if s.CommunityFinderRecruiting {
		cfPublic = 1
	}

	builtURI := fmt.Sprintf(uriFormat, name, s.DC, s.ActiveMembers, cfPublic, s.Order)
	return builtURI
}
