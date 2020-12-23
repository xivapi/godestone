package selectors

import (
	"encoding/json"

	"github.com/karashiiro/godestone/pack"
)

// FreeCompanyBasicSelectors contains the CSS selectors for the basic information on the Free Company page.
type FreeCompanyBasicSelectors struct {
	ActiveState       SelectorInfo `json:"ACTIVE_STATE"`
	ActiveMemberCount SelectorInfo `json:"ACTIVE_MEMBER_COUNT"`
	CrestLayers       struct {
		Bottom SelectorInfo `json:"BOTTOM"`
		Middle SelectorInfo `json:"MIDDLE"`
		Top    SelectorInfo `json:"TOP"`
	} `json:"CREST_LAYERS"`
	Estate struct {
		NoEstate SelectorInfo `json:"NO_ESTATE"`
		Greeting SelectorInfo `json:"GREETING"`
		Name     SelectorInfo `json:"NAME"`
		Plot     SelectorInfo `json:"PLOT"`
	} `json:"ESTATE"`
	Formed       SelectorInfo `json:"FORMED"`
	GrandCompany SelectorInfo `json:"GRAND_COMPANY"`
	ID           SelectorInfo `json:"ID"`
	Name         SelectorInfo `json:"NAME"`
	Rank         SelectorInfo `json:"RANK"`
	Ranking      struct {
		Monthly SelectorInfo `json:"MONTHLY"`
		Weekly  SelectorInfo `json:"WEEKLY"`
	} `json:"RANKING"`
	Recruitment SelectorInfo `json:"RECRUITMENT"`
	Server      SelectorInfo `json:"SERVER"`
	Slogan      SelectorInfo `json:"SLOGAN"`
	Tag         SelectorInfo `json:"TAG"`
}

// FreeCompanyMemberSelectors contains the CSS selectors for the member list on the Free Company page.
type FreeCompanyMemberSelectors struct {
	EntriesContainer SelectorInfo `json:"ENTRIES_CONTAINER"`
	Entry            struct {
		Root     SelectorInfo `json:"ROOT"`
		Avatar   SelectorInfo `json:"AVATAR"`
		ID       SelectorInfo `json:"ID"`
		Name     SelectorInfo `json:"NAME"`
		Rank     SelectorInfo `json:"RANK"`
		RankIcon SelectorInfo `json:"RANK_ICON"`
		Server   SelectorInfo `json:"SERVER"`
	} `json:"ENTRY"`
}

// FreeCompanyFocusSelectors contains the CSS selectors for a single focus on the Free Company page.
type FreeCompanyFocusSelectors struct {
	Name   SelectorInfo `json:"NAME"`
	Icon   SelectorInfo `json:"ICON"`
	Status SelectorInfo `json:"STATUS"`
}

// FreeCompanyFocusListSelectors contains the CSS selectors for the focus list on the Free Company page.
type FreeCompanyFocusListSelectors struct {
	NotSpecified SelectorInfo              `json:"NOT_SPECIFIED"`
	RolePlaying  FreeCompanyFocusSelectors `json:"RP"`
	Leveling     FreeCompanyFocusSelectors `json:"LEVELING"`
	Casual       FreeCompanyFocusSelectors `json:"CASUAL"`
	Hardcore     FreeCompanyFocusSelectors `json:"HARDCORE"`
	Dungeons     FreeCompanyFocusSelectors `json:"DUNGEONS"`
	Guildhests   FreeCompanyFocusSelectors `json:"GUILDHESTS"`
	Trials       FreeCompanyFocusSelectors `json:"TRIALS"`
	Raids        FreeCompanyFocusSelectors `json:"RAIDS"`
	PVP          FreeCompanyFocusSelectors `json:"PVP"`
}

// FreeCompanyAlignmentSelectors contains the CSS selectors for a single Grand Company Alignment on the Free Company page.
type FreeCompanyAlignmentSelectors struct {
	Name     SelectorInfo `json:"NAME"`
	Progress SelectorInfo `json:"PROGRESS"`
	Rank     SelectorInfo `json:"RANK"`
}

// FreeCompanyAlignmentListSelectors containes the CSS selectors for all of an FC's Grand Company aLignments.
type FreeCompanyAlignmentListSelectors struct {
	Maelstrom FreeCompanyAlignmentSelectors `json:"MAELSTROM"`
	Adders    FreeCompanyAlignmentSelectors `json:"ADDERS"`
	Flames    FreeCompanyAlignmentSelectors `json:"FLAMES"`
}

// FreeCompanySeekingSelectors contains the CSS selectors for a single seeking status on the Free Company page.
type FreeCompanySeekingSelectors struct {
	Name   SelectorInfo `json:"NAME"`
	Icon   SelectorInfo `json:"ICON"`
	Status SelectorInfo `json:"STATUS"`
}

// FreeCompanySeekingListSelectors contains the CSS selectors for the seeking status list on the Free Company page.
type FreeCompanySeekingListSelectors struct {
	NotSpecified SelectorInfo                `json:"NOT_SPECIFIED"`
	Tank         FreeCompanySeekingSelectors `json:"TANK"`
	Healer       FreeCompanySeekingSelectors `json:"HEALER"`
	DPS          FreeCompanySeekingSelectors `json:"DPS"`
	Crafter      FreeCompanySeekingSelectors `json:"CRAFTER"`
	Gatherer     FreeCompanySeekingSelectors `json:"GATHERER"`
}

// FreeCompanySelectors contains the CSS selectors for the Free Company page.
type FreeCompanySelectors struct {
	Basic      *FreeCompanyBasicSelectors
	Members    *FreeCompanyMemberSelectors
	Focuses    *FreeCompanyFocusListSelectors
	Reputation *FreeCompanyAlignmentListSelectors
	Seeking    *FreeCompanySeekingListSelectors
}

// LoadFreeCompanySelectors loads the CSS selectors for the Free Company page.
func LoadFreeCompanySelectors() (*FreeCompanySelectors, error) {
	basicBytes, err := pack.Asset("freecompany/freecompany.json")
	if err != nil {
		return nil, err
	}
	basicSelectors := FreeCompanyBasicSelectors{}
	json.Unmarshal(basicBytes, &basicSelectors)

	membersBytes, err := pack.Asset("freecompany/members.json")
	if err != nil {
		return nil, err
	}
	membersSelectors := FreeCompanyMemberSelectors{}
	json.Unmarshal(membersBytes, &membersSelectors)

	focusBytes, err := pack.Asset("freecompany/focus.json")
	if err != nil {
		return nil, err
	}
	focusSelectors := FreeCompanyFocusListSelectors{}
	json.Unmarshal(focusBytes, &focusSelectors)

	repBytes, err := pack.Asset("freecompany/reputation.json")
	if err != nil {
		return nil, err
	}
	repSelectors := FreeCompanyAlignmentListSelectors{}
	json.Unmarshal(repBytes, &repSelectors)

	seekBytes, err := pack.Asset("freecompany/seeking.json")
	if err != nil {
		return nil, err
	}
	seekSelectors := FreeCompanySeekingListSelectors{}
	json.Unmarshal(seekBytes, &seekSelectors)

	return &FreeCompanySelectors{
		Basic:      &basicSelectors,
		Members:    &membersSelectors,
		Focuses:    &focusSelectors,
		Reputation: &repSelectors,
		Seeking:    &seekSelectors,
	}, nil
}
