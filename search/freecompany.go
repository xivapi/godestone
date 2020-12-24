package search

import (
	"fmt"
	"strings"
)

// FreeCompanyActiveState is the active state of an FC.
type FreeCompanyActiveState string

// Active state for an FC.
const (
	FCActiveNotSpecified FreeCompanyActiveState = "Not specified"
	FCActiveWeekdaysOnly FreeCompanyActiveState = "Weekdays Only"
	FCActiveWeekendsOnly FreeCompanyActiveState = "Weekends Only"
	FCActiveAlways       FreeCompanyActiveState = "Always"
)

// FreeCompanyRecruitingState is the recruiting state of an FC.
type FreeCompanyRecruitingState string

// Recruiting state for an FC.
const (
	FCRecruitmentClosed FreeCompanyRecruitingState = "Closed"
	FCRecruitmentOpen   FreeCompanyRecruitingState = "Open"
)

// FreeCompanyFocus is an FC's focus.
type FreeCompanyFocus string

// Free Company Focus.
const (
	FCFocusRolePlaying FreeCompanyFocus = "Role-playing"
	FCFocusLeveling    FreeCompanyFocus = "Leveling"
	FCFocusCasual      FreeCompanyFocus = "Casual"
	FCFocusHardcore    FreeCompanyFocus = "Hardcore"
	FCFocusDungeons    FreeCompanyFocus = "Dungeons"
	FCFocusGuildhests  FreeCompanyFocus = "Guildhests"
	FCFocusTrials      FreeCompanyFocus = "Trials"
	FCFocusRaids       FreeCompanyFocus = "Raids"
	FCFocusPvP         FreeCompanyFocus = "PvP"
)

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
