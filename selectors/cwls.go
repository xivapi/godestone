package selectors

import (
	"encoding/json"

	"github.com/karashiiro/godestone/pack/css"
)

// CWLSBasicSelectors contains the CSS selectors for the basic information on the CWLS page.
type CWLSBasicSelectors struct {
	Name SelectorInfo `json:"NAME"`
	DC   SelectorInfo `json:"DC"`
}

// CWLSMemberSelectors contains the CSS selectors for the member list on the CWLS page.
type CWLSMemberSelectors struct {
	Root  SelectorInfo `json:"ROOT"`
	Entry struct {
		Root              SelectorInfo `json:"ROOT"`
		Avatar            SelectorInfo `json:"AVATAR"`
		ID                SelectorInfo `json:"ID"`
		Name              SelectorInfo `json:"NAME"`
		Rank              SelectorInfo `json:"RANK"`
		RankIcon          SelectorInfo `json:"RANK_ICON"`
		LinkshellRank     SelectorInfo `json:"LINKSHELL_RANK"`
		LinkshellRankIcon SelectorInfo `json:"LINKSHELL_RANK_ICON"`
		Server            SelectorInfo `json:"SERVER"`
	} `json:"ENTRY"`
	ListNextButton SelectorInfo `json:"LIST_NEXT_BUTTON"`
}

// CWLSSelectors contains the CSS selectors for the CWLS page.
type CWLSSelectors struct {
	Basic   *CWLSBasicSelectors
	Members *CWLSMemberSelectors
}

// LoadCWLSSelectors loads the CSS selectors for the CWLS page.
func LoadCWLSSelectors() *CWLSSelectors {
	basicBytes, _ := css.Asset("cwls/cwls.json")
	basicSelectors := CWLSBasicSelectors{}
	json.Unmarshal(basicBytes, &basicSelectors)

	membersBytes, _ := css.Asset("cwls/members.json")
	membersSelectors := CWLSMemberSelectors{}
	json.Unmarshal(membersBytes, &membersSelectors)

	return &CWLSSelectors{
		Basic:   &basicSelectors,
		Members: &membersSelectors,
	}
}
