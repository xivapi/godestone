package selectors

import (
	"encoding/json"

	"github.com/karashiiro/godestone/pack"
)

// PVPTeamBasicSelectors contains the CSS selectors for the basic information on the PVP team page.
type PVPTeamBasicSelectors struct {
	Name        SelectorInfo `json:"NAME"`
	DC          SelectorInfo `json:"DC"`
	Formed      SelectorInfo `json:"FORMED"`
	CrestLayers struct {
		Bottom SelectorInfo `json:"BOTTOM"`
		Middle SelectorInfo `json:"MIDDLE"`
		Top    SelectorInfo `json:"TOP"`
	} `json:"CREST_LAYERS"`
}

// PVPTeamMemberSelectors contains the CSS selectors for the member list on the PVP team page.
type PVPTeamMemberSelectors struct {
	EntriesContainer SelectorInfo `json:"ENTRIES_CONTAINER"`
	Entry            struct {
		Root     SelectorInfo `json:"ROOT"`
		Avatar   SelectorInfo `json:"AVATAR"`
		ID       SelectorInfo `json:"ID"`
		Name     SelectorInfo `json:"NAME"`
		Matches  SelectorInfo `json:"MATCHES"`
		Rank     SelectorInfo `json:"RANK"`
		RankIcon SelectorInfo `json:"RANK_ICON"`
		Server   SelectorInfo `json:"SERVER"`
	} `json:"ENTRY"`
}

// PVPTeamSelectors contains the CSS selectors for the PVP team page.
type PVPTeamSelectors struct {
	Basic   *PVPTeamBasicSelectors
	Members *PVPTeamMemberSelectors
}

// LoadPVPTeamSelectors loads the CSS selectors for the PVP team page.
func LoadPVPTeamSelectors() (*PVPTeamSelectors, error) {
	basicBytes, err := pack.Asset("pvpteam/pvpteam.json")
	if err != nil {
		return nil, err
	}
	basicSelectors := PVPTeamBasicSelectors{}
	json.Unmarshal(basicBytes, &basicSelectors)

	membersBytes, err := pack.Asset("pvpteam/members.json")
	if err != nil {
		return nil, err
	}
	membersSelectors := PVPTeamMemberSelectors{}
	json.Unmarshal(membersBytes, &membersSelectors)

	return &PVPTeamSelectors{
		Basic:   &basicSelectors,
		Members: &membersSelectors,
	}, nil
}
