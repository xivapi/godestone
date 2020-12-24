package models

import (
	"time"

	"github.com/karashiiro/godestone/data/gcrank"
	"github.com/karashiiro/godestone/data/reputation"
	"github.com/karashiiro/godestone/data/role"
	"github.com/karashiiro/godestone/search"

	"github.com/karashiiro/godestone/data/grandcompany"
)

// FreeCompanyFocusInfo represents a particular FC's intentions for a focus.
type FreeCompanyFocusInfo struct {
	Icon   string
	Kind   search.FreeCompanyFocus
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

// FreeCompanyMember represents information about a Free Company member.
type FreeCompanyMember struct {
	Error error

	Avatar   string
	ID       uint32
	Name     string
	Rank     gcrank.GCRank
	RankIcon string
	World    string
	DC       string
}

// FreeCompany represents all of the basic information about an FC.
type FreeCompany struct {
	Active            search.FreeCompanyActiveState
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
	Recruitment       search.FreeCompanyRecruitingState
	Reputation        []*FreeCompanyReputation
	Seeking           []*FreeCompanySeekingInfo
	Slogan            string
	Tag               string
	World             string
}

// FreeCompanySearchResult represents all of the searchable information about an FC.
type FreeCompanySearchResult struct {
	Error error

	Active        search.FreeCompanyActiveState
	ActiveMembers uint32
	CrestLayers   *CrestLayers
	DC            string
	Estate        string
	Formed        time.Time
	GrandCompany  grandcompany.GrandCompany
	ID            string
	Name          string
	Recruitment   search.FreeCompanyRecruitingState
	World         string
}
