package models

import (
	"time"

	"github.com/karashiiro/godestone/data/reputation"
	"github.com/karashiiro/godestone/data/role"

	"github.com/karashiiro/godestone/data/grandcompany"
)

// FreeCompanyActiveState is the active state of an FC.
type FreeCompanyActiveState string

// Active state for an FC.
const (
	FCActiveNotSpecified FreeCompanyActiveState = "Not specified"
	FCActiveAlways       FreeCompanyActiveState = "Always"
)

// FreeCompanyRecruitingState is the recruiting state of an FC.
type FreeCompanyRecruitingState string

// Recruiting state for an FC.
const (
	FCActiveClosed FreeCompanyRecruitingState = "Closed"
	FCActiveOpen   FreeCompanyRecruitingState = "Open"
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

// FreeCompanyFocusInfo represents a particular FC's intentions for a focus.
type FreeCompanyFocusInfo struct {
	Icon   string
	Kind   FreeCompanyFocus
	Status bool
}

// FreeCompanyRanking represents a particular FC's periodic rankings.
type FreeCompanyRanking struct {
	Monthly uint32
	Weekly  uint32
}

// FreeCompanyReputation represents an FC's alignment with each Grand Company.
type FreeCompanyReputation struct {
	GrandCompany grandcompany.GrandCompany
	Progress     uint8
	Rank         reputation.Reputation
}

// FreeCompanySeekingInfo represents a particular FC's intentions for a recruit roles.
type FreeCompanySeekingInfo struct {
	Icon   string
	Kind   role.Role
	Status bool
}

// FreeCompany represents all of the basic information about an FC.
type FreeCompany struct {
	Active            FreeCompanyActiveState
	ActiveMemberCount uint32
	CrestLayers       *CrestLayers
	DC                string
	Estate            *Estate
	Focus             []*FreeCompanyFocusInfo
	Formed            time.Time
	GrandCompany      grandcompany.GrandCompany
	ID                string
	Name              string
	ParseDate         time.Time
	Rank              uint8
	Ranking           *FreeCompanyRanking
	Recruitment       FreeCompanyRecruitingState
	Reputation        []*FreeCompanyReputation
	Seeking           []*FreeCompanySeekingInfo
	Slogan            string
	Tag               string
	World             string
}
