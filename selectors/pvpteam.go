package selectors

import (
	"encoding/json"

	"github.com/xivapi/godestone/pack/css"
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
	Root  SelectorInfo `json:"ROOT"`
	Entry struct {
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
func LoadPVPTeamSelectors() *PVPTeamSelectors {
	basicBytes, _ := css.Asset("pvpteam/pvpteam.json")
	basicSelectors := PVPTeamBasicSelectors{}
	json.Unmarshal(basicBytes, &basicSelectors)

	membersBytes, _ := css.Asset("pvpteam/members.json")
	membersSelectors := PVPTeamMemberSelectors{}
	json.Unmarshal(membersBytes, &membersSelectors)

	return &PVPTeamSelectors{
		Basic:   &basicSelectors,
		Members: &membersSelectors,
	}
}
