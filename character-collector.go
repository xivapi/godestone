package godestone

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/gocolly/colly/v2"
	"github.com/karashiiro/godestone/data/deity"
	"github.com/karashiiro/godestone/data/gcrank"
	"github.com/karashiiro/godestone/data/gender"
	"github.com/karashiiro/godestone/data/grandcompany"
	"github.com/karashiiro/godestone/data/race"
	"github.com/karashiiro/godestone/data/town"
	"github.com/karashiiro/godestone/data/tribe"
	"github.com/karashiiro/godestone/models"
	"github.com/karashiiro/godestone/selectors"
)

func (s *Scraper) makeCharCollector(charData *models.Character) *colly.Collector {
	c := colly.NewCollector()
	c.UserAgent = s.meta.UserAgentDesktop
	c.IgnoreRobotsTxt = true

	charSelectors := s.profileSelectors.Character

	c.OnHTML(charSelectors.Avatar.Selector, func(e *colly.HTMLElement) {
		charData.Avatar = e.Attr("src")
	})

	c.OnHTML(charSelectors.Bio.Selector, func(e *colly.HTMLElement) {
		charData.Bio = e.Text
	})

	fcIDRegex := regexp.MustCompile(charSelectors.FreeCompany.Name.Regex)
	c.OnHTML(charSelectors.FreeCompany.Name.Selector, func(e *colly.HTMLElement) {
		matches := fcIDRegex.FindStringSubmatch(e.Attr("href"))
		if matches != nil {
			/*
				This could be parsed to a uint64, but I don't know what SE's theoretical cap on Free Companies is and I'd
				rather this not break in a decade. It's harmless to keep it as a string, anyways, since it needs to be
				onverted to one to do a Lodestone lookup with it and anyone who wants it as a uint64 can just convert it themselves.
			*/
			charData.FreeCompanyID = matches[1]
			charData.FreeCompanyName = e.Text
		}
	})

	c.OnHTML(charSelectors.GrandCompany.Selector, func(e *colly.HTMLElement) {
		gcRawInfo := strings.Split(e.Text, "/")
		gcName := gcRawInfo[0][0 : len(gcRawInfo[0])-1]
		gcRankNameParts := strings.Split(gcRawInfo[1][1:], " ")
		gcRank := gcRankNameParts[len(gcRankNameParts)-1]

		gcID := grandcompany.Parse(gcName)
		gcRankID := gcrank.Parse(gcRank)

		gc := models.GrandCompanyInfo{NameID: gcID, RankID: gcRankID}
		charData.GrandCompany = &gc
	})

	c.OnHTML(charSelectors.GuardianDeity.Selector, func(e *colly.HTMLElement) {
		charData.GuardianDeity = deity.Parse(e.Text)
	})

	c.OnHTML(charSelectors.Name.Selector, func(e *colly.HTMLElement) {
		charData.Name = e.Text
	})

	c.OnHTML(charSelectors.Nameday.Selector, func(e *colly.HTMLElement) {
		charData.Nameday = e.Text
	})

	c.OnHTML(charSelectors.Portrait.Selector, func(e *colly.HTMLElement) {
		charData.Portrait = e.Attr("src")
	})

	pvpTeamIDRegex := regexp.MustCompile(charSelectors.PvPTeam.Name.Regex)
	c.OnHTML(charSelectors.PvPTeam.Name.Selector, func(e *colly.HTMLElement) {
		matches := pvpTeamIDRegex.FindStringSubmatch(e.Attr("href"))
		if matches != nil {
			charData.PvPTeamID = matches[1]
		}
	})

	raceClanGenderRegex := regexp.MustCompile(charSelectors.RaceClanGender.Regex)
	c.OnHTML(charSelectors.RaceClanGender.Selector, func(e *colly.HTMLElement) {
		rawText, err := e.DOM.Html()
		if err != nil {
			return
		}

		matches := raceClanGenderRegex.FindStringSubmatch(rawText)
		if matches != nil {
			charData.Race = race.Parse(matches[1])
			charData.Tribe = tribe.Parse(matches[2])
			charData.Gender = gender.Parse(matches[3])
		}
	})

	c.OnHTML(charSelectors.Server.Selector, func(e *colly.HTMLElement) {
		server := e.Text
		serverSplit := strings.Split(server, "(")
		world := serverSplit[0][0 : len(serverSplit[0])-2]
		dc := serverSplit[1][0 : len(serverSplit[1])-1]

		charData.DC = dc
		charData.Server = world
	})

	c.OnHTML(charSelectors.Title.Selector, func(e *colly.HTMLElement) {
		// TODO
		charData.Title = 0
		charData.TitleTop = false
	})

	c.OnHTML(charSelectors.Town.Selector, func(e *colly.HTMLElement) {
		charData.Town = town.Parse(e.Text)
	})

	charData.GearSet = &models.GearSet{}

	partRefs := &models.GearItemBuild{
		Body:      &models.GearItem{},
		Bracelets: &models.GearItem{},
		Earrings:  &models.GearItem{},
		Feet:      &models.GearItem{},
		Hands:     &models.GearItem{},
		Head:      &models.GearItem{},
		Legs:      &models.GearItem{},
		MainHand:  &models.GearItem{},
		Necklace:  &models.GearItem{},
		OffHand:   &models.GearItem{},
		Ring1:     &models.GearItem{},
		Ring2:     &models.GearItem{},
		SoulCrystal: &models.GearItem{
			Materia: make([]uint32, 0),
		},
		Waist: &models.GearItem{},
	}

	charData.GearSet.Gear = partRefs
	partSelectors := s.profileSelectors.GearSet
	parts := map[*models.GearItem]*selectors.GearSelectors{
		partRefs.MainHand:  &partSelectors.MainHand,
		partRefs.OffHand:   &partSelectors.OffHand,
		partRefs.Head:      &partSelectors.Head,
		partRefs.Body:      &partSelectors.Body,
		partRefs.Hands:     &partSelectors.Hands,
		partRefs.Waist:     &partSelectors.Waist,
		partRefs.Legs:      &partSelectors.Legs,
		partRefs.Feet:      &partSelectors.Feet,
		partRefs.Earrings:  &partSelectors.Earrings,
		partRefs.Necklace:  &partSelectors.Necklace,
		partRefs.Bracelets: &partSelectors.Bracelets,
		partRefs.Ring1:     &partSelectors.Ring1,
		partRefs.Ring2:     &partSelectors.Ring2,
	}

	for partRef, partSelector := range parts {
		// Closures are fun
		currRef := partRef
		currSelector := partSelector

		currRef.Materia = make([]uint32, 0)

		c.OnHTML(currSelector.CreatorName.Selector, func(e *colly.HTMLElement) {
			currRef.Creator = e.Text
		})
		c.OnHTML(currSelector.Stain.Selector, func(e *colly.HTMLElement) {
			currRef.Dye = 0
		})
		c.OnHTML(currSelector.Name.Selector, func(e *colly.HTMLElement) {
			currRef.ID = 0
		})

		materiaCallback := func(e *colly.HTMLElement) {
			currRef.Materia = append(currRef.Materia, 0)
		}
		c.OnHTML(currSelector.Materia1.Selector, materiaCallback)
		c.OnHTML(currSelector.Materia2.Selector, materiaCallback)
		c.OnHTML(currSelector.Materia3.Selector, materiaCallback)
		c.OnHTML(currSelector.Materia4.Selector, materiaCallback)
		c.OnHTML(currSelector.Materia5.Selector, materiaCallback)

		c.OnHTML(currSelector.MirageName.Selector, func(e *colly.HTMLElement) {
			currRef.Mirage = 0
		})
	}

	c.OnHTML(partSelectors.SoulCrystal.Name.Selector, func(e *colly.HTMLElement) {
		partRefs.SoulCrystal.ID = 0
	})

	activeClassJobLevelRegex := regexp.MustCompile(charSelectors.ActiveClassJobLevel.Regex)
	c.OnHTML(charSelectors.ActiveClassJobLevel.Selector, func(e *colly.HTMLElement) {
		matches := activeClassJobLevelRegex.FindStringSubmatch(e.Text)
		if matches != nil {
			level, err := strconv.ParseUint(matches[1], 10, 8)
			if err != nil {
				charData.GearSet.Level = uint8(level)
			}
		}
	})

	return c
}
