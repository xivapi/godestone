package godestone

import (
	"fmt"
	"strings"
)

// CharacterSearchOrder represents the search result ordering of a Lodestone character search.
type CharacterSearchOrder uint8

// Search ordering for character searches.
const (
	OrderCharaNameAToZ CharacterSearchOrder = iota + 1
	OrderCharaNameZToA
	OrderCharaWorldAtoZ
	OrderCharaWorldZtoA
	OrderCharaLevelDescending
	OrderCharaLevelAscending
)

// CharacterOptions defines extra search information that can help to narrow down a search.
type CharacterOptions struct {
	Name         string
	World        string
	DC           string
	SearchLang   SearchLang
	GrandCompany string
	Race         string
	Tribe        string
	Order        CharacterSearchOrder
}

// BuildURI returns a constructed URI for the provided search options.
func (s *CharacterOptions) BuildURI(
	scraper *Scraper,
	lang string,
) string {
	uriFormat := "https://%s.finalfantasyxiv.com/lodestone/character/?q=%s&worldname=%s&classjob=%s&order=%d"

	name := strings.Replace(s.Name, " ", "%20", -1)

	worldDC := parseWorldDC(s.World, s.DC)

	if s.SearchLang == NoneLang || s.SearchLang&SearchJA != 0 {
		uriFormat += "&blog_lang=ja"
	}
	if s.SearchLang == NoneLang || s.SearchLang&SearchEN != 0 {
		uriFormat += "&blog_lang=en"
	}
	if s.SearchLang == NoneLang || s.SearchLang&SearchDE != 0 {
		uriFormat += "&blog_lang=de"
	}
	if s.SearchLang == NoneLang || s.SearchLang&SearchFR != 0 {
		uriFormat += "&blog_lang=fr"
	}

	if s.Tribe != "" || s.Race != "" {
		raceTribe := ""
		if s.Tribe != "" {
			t := scraper.dataProvider.Tribe(s.Tribe)
			raceTribe = fmt.Sprintf("tribe_%d", t.ID)
		} else if s.Race != "" {
			r := scraper.dataProvider.Race(s.Race)
			raceTribe = fmt.Sprintf("race_%d", r.ID)
		}
		uriFormat += fmt.Sprintf("&race_tribe=%s", raceTribe)
	}

	if s.GrandCompany != "" {
		gc := scraper.dataProvider.GrandCompany(s.GrandCompany)
		uriFormat += fmt.Sprintf("&gcid=%d", gc.ID)
	}

	builtURI := fmt.Sprintf(uriFormat, lang, name, worldDC, "", s.Order)
	return builtURI
}

// FreeCompanySearchOrder represents the search result ordering of a Lodestone Free Company search.
type FreeCompanySearchOrder uint8

// Search ordering for Free Company searches.
const (
	OrderFCNameAToZ FreeCompanySearchOrder = iota + 1
	OrderFCNameZToA
	OrderFCMembershipHighToLow
	OrderFCMembershipLowToHigh
	OrderFCDateFoundedDescending
	OrderFCDateFoundedAscending
)

// FreeCompanyHousingStatus represents the housing status of a Free Company for the purpose of searches.
type FreeCompanyHousingStatus uint8

// Housing status for Free Company searches.
const (
	FCHousingAll FreeCompanyHousingStatus = iota
	FCHousingNoEstateOrPlot
	FCHousingPlotOnly
	FCHousingEstateBuilt
)

// FreeCompanyOptions defines extra search information that can help to narrow down a Free Company search.
type FreeCompanyOptions struct {
	Name                      string
	World                     string
	DC                        string
	ActiveTime                FreeCompanyActiveState
	Recruitment               FreeCompanyRecruitingState
	Order                     FreeCompanySearchOrder
	HousingStatus             FreeCompanyHousingStatus
	ActiveMembers             ActiveMemberRange
	CommunityFinderRecruiting bool
}

// BuildURI returns a constructed URI for the provided search options.
func (s *FreeCompanyOptions) BuildURI(lang string) string {
	uriFormat := "https://%s.finalfantasyxiv.com/lodestone/freecompany/?q=%s&worldname=%s&character_count=%s&cf_public=%d&activetime=%s&join=%s&house=%s&order=%d"

	name := strings.Replace(s.Name, " ", "%20", -1)

	worldDC := parseWorldDC(s.World, s.DC)

	cfPublic := 0
	if s.CommunityFinderRecruiting {
		cfPublic = 1
	}

	join := ""
	if s.Recruitment == FCRecruitmentOpen {
		join = "1"
	} else if s.Recruitment == FCRecruitmentClosed {
		join = "0"
	}

	active := ""
	if s.ActiveTime == FCActiveWeekdaysOnly {
		active = "1"
	} else if s.ActiveTime == FCActiveWeekendsOnly {
		active = "2"
	}

	housingStatus := ""
	if s.HousingStatus != FCHousingAll {
		housingStatus = fmt.Sprint(s.HousingStatus)
	}

	builtURI := fmt.Sprintf(uriFormat, lang, name, worldDC, s.ActiveMembers, cfPublic, active, join, housingStatus, s.Order)
	return builtURI
}

// ActiveMemberRange represents the active member range filter of a search.
type ActiveMemberRange string

// Active member range for searches.
const (
	OneToTen         ActiveMemberRange = "1-10"
	ElevenToThirty   ActiveMemberRange = "11-30"
	ThirtyOneToFifty ActiveMemberRange = "31-50"
	FiftyOnePlus     ActiveMemberRange = "51-"
)

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
func (s *LinkshellOptions) BuildURI(lang string) string {
	uriFormat := "https://%s.finalfantasyxiv.com/lodestone/linkshell/?q=%s&worldname=%s&character_count=%s&cf_public=%d&order=%d"

	name := strings.Replace(s.Name, " ", "%20", -1)

	worldDC := parseWorldDC(s.World, s.DC)

	cfPublic := 0
	if s.CommunityFinderRecruiting {
		cfPublic = 1
	}

	builtURI := fmt.Sprintf(uriFormat, lang, name, worldDC, s.ActiveMembers, cfPublic, s.Order)
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
func (s *CWLSOptions) BuildURI(lang string) string {
	uriFormat := "https://%s.finalfantasyxiv.com/lodestone/crossworld_linkshell/?q=%s&dcname=%s&character_count=%s&cf_public=%d&order=%d"

	name := strings.Replace(s.Name, " ", "%20", -1)

	cfPublic := 0
	if s.CommunityFinderRecruiting {
		cfPublic = 1
	}

	builtURI := fmt.Sprintf(uriFormat, lang, name, s.DC, s.ActiveMembers, cfPublic, s.Order)
	return builtURI
}

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
