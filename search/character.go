package search

import (
	"fmt"
	"strings"

	"github.com/karashiiro/godestone/data/grandcompany"
	"github.com/karashiiro/godestone/data/race"
	"github.com/karashiiro/godestone/data/tribe"
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
	Lang         Lang
	GrandCompany grandcompany.GrandCompany
	Race         race.Race
	Tribe        tribe.Tribe
	Order        CharacterSearchOrder
}

// BuildURI returns a constructed URI for the provided search options.
func (s *CharacterOptions) BuildURI() string {
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
