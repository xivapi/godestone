package godestone

import (
	"fmt"
	"strings"

	"github.com/karashiiro/godestone/data/grandcompany"
	"github.com/karashiiro/godestone/data/race"
	"github.com/karashiiro/godestone/data/tribe"
)

// LinkshellSearchOrder represents the search result ordering of a Lodestone CWLS search.
type LinkshellSearchOrder uint8

// Search ordering for linkshell and CWLS searches.
const (
	OrderLinkshellNameAToZ            LinkshellSearchOrder = 1
	OrderLinkshellNameZToA            LinkshellSearchOrder = 2
	OrderLinkshellMembershipHighToLow LinkshellSearchOrder = 3
	OrderLinkshellMembershipLowToHigh LinkshellSearchOrder = 4
)

// LinkshellActiveMemberRange represents the active member range filter of a Lodestone CWLS search.
type LinkshellActiveMemberRange string

// Active member range for linkshell and CWLS searches.
const (
	OneToTen         LinkshellActiveMemberRange = "1-10"
	ElevenToThirty   LinkshellActiveMemberRange = "11-30"
	ThirtyOneToFifty LinkshellActiveMemberRange = "31-50"
	FiftyOnePlus     LinkshellActiveMemberRange = "51-"
)

// SearchLinkshellOptions defines extra search information that can help to narrow down a linkshell search.
type SearchLinkshellOptions struct {
	Name                      string
	World                     string
	DC                        string
	Order                     LinkshellSearchOrder
	ActiveMembers             LinkshellActiveMemberRange
	CommunityFinderRecruiting bool
}

func (s *SearchLinkshellOptions) buildURI() string {
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

// SearchCWLSOptions defines extra search information that can help to narrow down a CWLS search.
type SearchCWLSOptions struct {
	Name                      string
	DC                        string
	Order                     LinkshellSearchOrder
	ActiveMembers             LinkshellActiveMemberRange
	CommunityFinderRecruiting bool
}

func (s *SearchCWLSOptions) buildURI() string {
	uriFormat := "https://na.finalfantasyxiv.com/lodestone/crossworld_linkshell/?q=%s&dcname=%s&character_count=%s&cf_public=%d&order=%d"

	name := strings.Replace(s.Name, " ", "%20", -1)

	cfPublic := 0
	if s.CommunityFinderRecruiting {
		cfPublic = 1
	}

	builtURI := fmt.Sprintf(uriFormat, name, s.DC, s.ActiveMembers, cfPublic, s.Order)
	return builtURI
}

// CharacterSearchOrder represents the search result ordering of a Lodestone character search.
type CharacterSearchOrder uint8

// Search ordering for character searches.
const (
	OrderCharaNameAToZ        CharacterSearchOrder = 1
	OrderCharaNameZToA        CharacterSearchOrder = 2
	OrderCharaWorldAtoZ       CharacterSearchOrder = 3
	OrderCharaWorldZtoA       CharacterSearchOrder = 4
	OrderCharaLevelDescending CharacterSearchOrder = 5
	OrderCharaLevelAscending  CharacterSearchOrder = 6
)

// SearchCharacterOptions defines extra search information that can help to narrow down a search.
type SearchCharacterOptions struct {
	Name         string
	World        string
	DC           string
	Lang         Lang
	GrandCompany grandcompany.GrandCompany
	Race         race.Race
	Tribe        tribe.Tribe
	Order        CharacterSearchOrder
}

func (s *SearchCharacterOptions) buildURI() string {
	uriFormat := "https://na.finalfantasyxiv.com/lodestone/character/?q=%s&worldname=%s&classjob=%s&order=%d"

	name := strings.Replace(s.Name, " ", "%20", -1)

	worldDC := parseWorldDC(s.World, s.DC)

	if s.Lang == NoneLang || s.Lang&JA != 0 {
		uriFormat += "&blog_lang=ja"
	}
	if s.Lang == NoneLang || s.Lang&EN != 0 {
		uriFormat += "&blog_lang=en"
	}
	if s.Lang == NoneLang || s.Lang&DE != 0 {
		uriFormat += "&blog_lang=de"
	}
	if s.Lang == NoneLang || s.Lang&FR != 0 {
		uriFormat += "&blog_lang=fr"
	}

	if s.Tribe != tribe.None || s.Race != race.None {
		raceTribe := ""
		if s.Tribe != tribe.None {
			raceTribe = fmt.Sprintf("tribe_%d", s.Tribe)
		} else if s.Race != race.None {
			raceTribe = fmt.Sprintf("race_%d", s.Race)
		}
		uriFormat += fmt.Sprintf("&race_tribe=%s", raceTribe)
	}

	if s.GrandCompany != grandcompany.None {
		uriFormat += fmt.Sprintf("&gcid=%d", s.GrandCompany)
	}

	builtURI := fmt.Sprintf(uriFormat, name, worldDC, "", s.Order)
	return builtURI
}

func parseWorldDC(world string, dc string) string {
	worldDC := dc
	if len(world) != 0 {
		worldDC = world
	} else {
		// DCs have the _dc_ prefix attached to them
		if len(worldDC) != 0 && !strings.HasPrefix(worldDC, "_dc_") {
			worldDC = "_dc_" + worldDC
		}
	}
	return worldDC
}
