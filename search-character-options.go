package godestone

import (
	"fmt"
	"strings"

	"github.com/karashiiro/godestone/data/grandcompany"
	"github.com/karashiiro/godestone/data/race"
	"github.com/karashiiro/godestone/data/tribe"
)

// SearchCharacterOptions defines extra search information that can help to narrow down a search.
type SearchCharacterOptions struct {
	Name         string
	World        string
	DC           string
	Lang         Lang
	GrandCompany grandcompany.GrandCompany
	Race         race.Race
	Tribe        tribe.Tribe
	Order        SearchOrder
}

func (s *SearchCharacterOptions) buildURI() string {
	uriFormat := "https://na.finalfantasyxiv.com/lodestone/character/?q=%s&worldname=%s&classjob=%s&order=%d"

	name := strings.Replace(s.Name, " ", "%20", -1)

	worldDC := s.DC
	if len(s.World) != 0 {
		worldDC = s.World
	} else {
		// DCs have the _dc_ prefix attached to them
		if len(worldDC) != 0 && !strings.HasPrefix(worldDC, "_dc_") {
			worldDC = "_dc_" + worldDC
		}
	}

	if s.Lang == None || s.Lang&JA != 0 {
		uriFormat += "&blog_lang=ja"
	}
	if s.Lang == None || s.Lang&EN != 0 {
		uriFormat += "&blog_lang=en"
	}
	if s.Lang == None || s.Lang&DE != 0 {
		uriFormat += "&blog_lang=de"
	}
	if s.Lang == None || s.Lang&FR != 0 {
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
