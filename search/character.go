package search

import (
	"fmt"
	"strings"

	"github.com/karashiiro/godestone/pack/exports"
	lookups "github.com/karashiiro/godestone/table-lookups"
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
	GrandCompany string
	Race         string
	Tribe        string
	Order        CharacterSearchOrder
}

// BuildURI returns a constructed URI for the provided search options.
func (s *CharacterOptions) BuildURI(
	grandCompanyTable *exports.GrandCompanyTable,
	raceTable *exports.RaceTable,
	tribeTable *exports.TribeTable,
	lang string,
) string {
	uriFormat := "https://%s.finalfantasyxiv.com/lodestone/character/?q=%s&worldname=%s&classjob=%s&order=%d"

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

	if s.Tribe != "" || s.Race != "" {
		raceTribe := ""
		if s.Tribe != "" {
			t := lookups.TribeTableLookup(tribeTable, s.Tribe)
			raceTribe = fmt.Sprintf("tribe_%d", t.Id())
		} else if s.Race != "" {
			r := lookups.RaceTableLookup(raceTable, s.Race)
			raceTribe = fmt.Sprintf("race_%d", r.Id())
		}
		uriFormat += fmt.Sprintf("&race_tribe=%s", raceTribe)
	}

	if s.GrandCompany != "" {
		gc := lookups.GrandCompanyTableLookup(grandCompanyTable, s.GrandCompany)
		uriFormat += fmt.Sprintf("&gcid=%d", gc.Id())
	}

	builtURI := fmt.Sprintf(uriFormat, lang, name, worldDC, "", s.Order)
	return builtURI
}
