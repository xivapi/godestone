package selectors

import (
	"encoding/json"

	"github.com/karashiiro/godestone/pack"
)

// CharacterSearchSelectors contains the CSS selectors for the character search interface.
type CharacterSearchSelectors struct {
	EntriesContainer SelectorInfo `json:"ENTRIES_CONTAINER"`
	Entry            struct {
		Root     SelectorInfo `json:"ROOT"`
		Avatar   SelectorInfo `json:"AVATAR"`
		ID       SelectorInfo `json:"ID"`
		Lang     SelectorInfo `json:"LANG"`
		Name     SelectorInfo `json:"NAME"`
		Rank     SelectorInfo `json:"RANK"`
		RankIcon SelectorInfo `json:"RANK_ICON"`
		Server   SelectorInfo `json:"SERVER"`
	}
	ListNextButton SelectorInfo `json:"LIST_NEXT_BUTTON"`
	PageInfo       SelectorInfo `json:"PAGE_INFO"`
}

// CWLSSearchSelectors contains the CSS selectors for the character search interface.
type CWLSSearchSelectors struct {
	EntriesContainer SelectorInfo `json:"ENTRIES_CONTAINER"`
	Entry            struct {
		Root          SelectorInfo `json:"ROOT"`
		ID            SelectorInfo `json:"ID"`
		Name          SelectorInfo `json:"NAME"`
		DC            SelectorInfo `json:"DC"`
		ActiveMembers SelectorInfo `json:"ACTIVE_MEMBERS"`
	}
	ListNextButton SelectorInfo `json:"LIST_NEXT_BUTTON"`
	PageInfo       SelectorInfo `json:"PAGE_INFO"`
}

// SearchSelectors contains the CSS selectors for the search interface.
type SearchSelectors struct {
	Character *CharacterSearchSelectors
	CWLS      *CWLSSearchSelectors
}

// LoadSearchSelectors loads the CSS selectors for the search interface.
func LoadSearchSelectors() (*SearchSelectors, error) {
	charaBytes, err := pack.Asset("search/character.json")
	if err != nil {
		return nil, err
	}
	charaSearchSelectors := CharacterSearchSelectors{}
	json.Unmarshal(charaBytes, &charaSearchSelectors)

	cwlsBytes, err := pack.Asset("search/cwls.json")
	if err != nil {
		return nil, err
	}
	cwlsSearchSelectors := CWLSSearchSelectors{}
	json.Unmarshal(cwlsBytes, &cwlsSearchSelectors)

	return &SearchSelectors{
		Character: &charaSearchSelectors,
		CWLS:      &cwlsSearchSelectors,
	}, nil
}
