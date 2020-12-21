package godestone

import (
	"strconv"

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
		charData.Avatar = charSelectors.Avatar.Parse(e)[0]
	})

	c.OnHTML(charSelectors.Bio.Selector, func(e *colly.HTMLElement) {
		charData.Bio = charSelectors.Bio.Parse(e)[0]
	})

	c.OnHTML(charSelectors.FreeCompany.Name.Selector, func(e *colly.HTMLElement) {
		fcID := charSelectors.FreeCompany.Name.Parse(e)[0]
		if fcID != "" {
			/*
				This could be parsed to a uint64, but I don't know what SE's theoretical cap on Free Companies is and I'd
				rather this not break in a decade. It's harmless to keep it as a string, anyways, since it needs to be
				onverted to one to do a Lodestone lookup with it and anyone who wants it as a uint64 can just convert it themselves.
			*/
			charData.FreeCompanyID = fcID
			charData.FreeCompanyName = e.Text
		}
	})

	c.OnHTML(charSelectors.GrandCompany.Selector, func(e *colly.HTMLElement) {
		values := charSelectors.GrandCompany.Parse(e)

		gcName := grandcompany.Parse(values[0])
		gcRank := gcrank.Parse(values[1])

		gc := models.GrandCompanyInfo{NameID: gcName, RankID: gcRank}
		charData.GrandCompany = &gc
	})

	c.OnHTML(charSelectors.GuardianDeity.Selector, func(e *colly.HTMLElement) {
		charData.GuardianDeity = deity.Parse(e.Text)
	})

	c.OnHTML(charSelectors.Name.Selector, func(e *colly.HTMLElement) {
		charData.Name = charSelectors.Name.Parse(e)[0]
	})

	c.OnHTML(charSelectors.Nameday.Selector, func(e *colly.HTMLElement) {
		charData.Nameday = charSelectors.Nameday.Parse(e)[0]
	})

	c.OnHTML(charSelectors.Portrait.Selector, func(e *colly.HTMLElement) {
		charData.Portrait = charSelectors.Portrait.Parse(e)[0]
	})

	c.OnHTML(charSelectors.PvPTeam.Name.Selector, func(e *colly.HTMLElement) {
		charData.PvPTeamID = charSelectors.PvPTeam.Name.Parse(e)[0]
	})

	c.OnHTML(charSelectors.RaceClanGender.Selector, func(e *colly.HTMLElement) {
		values := charSelectors.RaceClanGender.ParseInnerHTML(e)

		charData.Race = race.Parse(values[0])
		charData.Tribe = tribe.Parse(values[1])
		charData.Gender = gender.Parse(values[2])
	})

	c.OnHTML(charSelectors.Server.Selector, func(e *colly.HTMLElement) {
		values := charSelectors.Server.Parse(e)

		charData.Server = values[0]
		charData.DC = values[1]
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
			currRef.Creator = currSelector.CreatorName.Parse(e)[0]
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

	c.OnHTML(charSelectors.ActiveClassJobLevel.Selector, func(e *colly.HTMLElement) {
		levelStr := charSelectors.ActiveClassJobLevel.Parse(e)[0]
		level, err := strconv.ParseUint(levelStr, 10, 8)
		if err != nil {
			charData.GearSet.Level = uint8(level)
		}
	})

	return c
}
