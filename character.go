package godestone

import (
	"strconv"
	"strings"

	"github.com/gocolly/colly/v2"
	"github.com/xivapi/godestone/v2/data/baseparam"
	"github.com/xivapi/godestone/v2/data/gcrank"
	"github.com/xivapi/godestone/v2/data/gender"
	"github.com/xivapi/godestone/v2/provider/models"
	"github.com/xivapi/godestone/v2/selectors"
)

func (s *Scraper) buildCharacterCollector(
	charData *Character,
) *colly.Collector {
	c := colly.NewCollector(
		colly.UserAgent(s.meta.UserAgentDesktop),
		colly.IgnoreRobotsTxt(),
		colly.Async(),
	)

	// BASIC DATA
	charSelectors := s.getProfileSelectors().Character

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

		gcName := values[0]
		gcRank := gcrank.Parse(values[1])

		gc := s.dataProvider.GrandCompany(gcName)

		charData.GrandCompanyInfo = &GrandCompanyInfo{GrandCompany: gc, RankID: gcRank}
	})

	charData.GuardianDeity = &IconedNamedEntity{}
	c.OnHTML(charSelectors.GuardianDeity.Name.Selector, func(e *colly.HTMLElement) {
		name := charSelectors.GuardianDeity.Name.Parse(e)[0]
		d := s.dataProvider.Deity(name)
		charData.GuardianDeity.NamedEntity = d
	})
	c.OnHTML(charSelectors.GuardianDeity.Icon.Selector, func(e *colly.HTMLElement) {
		charData.GuardianDeity.Icon = charSelectors.GuardianDeity.Icon.Parse(e)[0]
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

		// Miqo'te fix
		raceName := strings.ReplaceAll(values[0], "&#39;", "'")

		r := s.dataProvider.Race(raceName)
		charData.Race = r

		t := s.dataProvider.Tribe(values[1])
		charData.Tribe = t

		charData.Gender = gender.Parse(values[2])
	})

	c.OnHTML(charSelectors.Server.Selector, func(e *colly.HTMLElement) {
		values := charSelectors.Server.Parse(e)

		charData.World = values[0]
		charData.DC = values[1]
	})

	c.OnHTML(charSelectors.Title.Selector, func(e *colly.HTMLElement) {
		titleText := charSelectors.Title.Parse(e)[0]
		t := s.dataProvider.Title(titleText)

		if t != nil {
			charData.Title = &Title{
				TitleInternal: t,
			}
		} else {
			charData.Title = &Title{
				TitleInternal: &models.TitleInternal{
					GenderedEntity: &models.GenderedEntity{
						ID:   0,
						Name: titleText,

						NameMasculineEN: "",
						NameMasculineDE: "",
						NameMasculineFR: "",
						NameMasculineJA: "",
						NameFeminineEN:  "",
						NameFeminineDE:  "",
						NameFeminineFR:  "",
						NameFeminineJA:  "",
					},
					Prefix: false,
				},
			}
		}
	})

	charData.Town = &IconedNamedEntity{}
	c.OnHTML(charSelectors.Town.Name.Selector, func(e *colly.HTMLElement) {
		name := charSelectors.Town.Name.Parse(e)[0]
		t := s.dataProvider.Town(name)

		charData.Town.NamedEntity = t
	})
	c.OnHTML(charSelectors.Town.Icon.Selector, func(e *colly.HTMLElement) {
		charData.Town.Icon = charSelectors.Town.Icon.Parse(e)[0]
	})

	charData.GearSet = &GearSet{}

	// ATTRIBUTES
	attributeSelectors := s.getProfileSelectors().Attributes
	charData.GearSet.Attributes = map[baseparam.BaseParam]uint32{}

	attributesMap := map[baseparam.BaseParam]*selectors.SelectorInfo{
		baseparam.Strength:            &attributeSelectors.Strength,
		baseparam.Dexterity:           &attributeSelectors.Dexterity,
		baseparam.Vitality:            &attributeSelectors.Vitality,
		baseparam.Intelligence:        &attributeSelectors.Intelligence,
		baseparam.Mind:                &attributeSelectors.Mind,
		baseparam.CriticalHit:         &attributeSelectors.CriticalHitRate,
		baseparam.Determination:       &attributeSelectors.Determination,
		baseparam.DirectHitRate:       &attributeSelectors.DirectHitRate,
		baseparam.Defense:             &attributeSelectors.Defense,
		baseparam.MagicDefense:        &attributeSelectors.MagicDefense,
		baseparam.AttackPower:         &attributeSelectors.AttackPower,
		baseparam.SkillSpeed:          &attributeSelectors.SkillSpeed,
		baseparam.AttackMagicPotency:  &attributeSelectors.AttackMagicPotency,
		baseparam.HealingMagicPotency: &attributeSelectors.HealingMagicPotency,
		baseparam.SpellSpeed:          &attributeSelectors.SpellSpeed,
		baseparam.Tenacity:            &attributeSelectors.Tenacity,
		baseparam.Piety:               &attributeSelectors.Piety,
		baseparam.HP:                  &attributeSelectors.HP,
		baseparam.MP:                  &attributeSelectors.MPGPCP,
		baseparam.GP:                  &attributeSelectors.MPGPCP,
		baseparam.CP:                  &attributeSelectors.MPGPCP,
	}

	resourceAttr := "" // MP, GP, or CP
	c.OnHTML(attributeSelectors.MPGPCPParameterName.Selector, func(e *colly.HTMLElement) {
		resourceAttr = attributeSelectors.MPGPCPParameterName.Parse(e)[0]
	})

	for attribute, selector := range attributesMap {
		currAttribute := attribute
		currSelector := selector

		c.OnHTML(currSelector.Selector, func(e *colly.HTMLElement) {
			valStr := currSelector.Parse(e)[0]
			val, err := strconv.ParseUint(valStr, 10, 32)
			if err == nil {
				if currAttribute == baseparam.MP || currAttribute == baseparam.GP || currAttribute == baseparam.CP {
					switch resourceAttr {
					case "MP":
						charData.GearSet.Attributes[baseparam.MP] = uint32(val)
					case "GP":
						charData.GearSet.Attributes[baseparam.GP] = uint32(val)
					case "CP":
						charData.GearSet.Attributes[baseparam.CP] = uint32(val)
					}
				} else {
					charData.GearSet.Attributes[currAttribute] = uint32(val)
				}
			}
		})
	}

	// GEAR PIECES
	partRefs := &GearItemBuild{
		Body:      &GearItem{},
		Bracelets: &GearItem{},
		Earrings:  &GearItem{},
		Feet:      &GearItem{},
		Hands:     &GearItem{},
		Head:      &GearItem{},
		Legs:      &GearItem{},
		MainHand:  &GearItem{},
		Necklace:  &GearItem{},
		OffHand:   &GearItem{},
		Ring1:     &GearItem{},
		Ring2:     &GearItem{},
		SoulCrystal: &GearItem{
			Materia: make([]uint32, 0),
		},
		Waist: &GearItem{},
	}

	charData.GearSet.Gear = partRefs
	partSelectors := s.getProfileSelectors().GearSet
	parts := map[*GearItem]*selectors.GearSelectors{
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
		currRef := partRef
		currSelector := partSelector

		currRef.Materia = make([]uint32, 0)

		c.OnHTML(currSelector.CreatorName.Selector, func(e *colly.HTMLElement) {
			currRef.Creator = currSelector.CreatorName.Parse(e)[0]
		})
		c.OnHTML(currSelector.Stain.Selector, func(e *colly.HTMLElement) {
			name := currSelector.Stain.Parse(e)[0]
			item := s.dataProvider.Item(name)
			if item != nil {
				currRef.Dye = item.ID
			}
		})
		c.OnHTML(currSelector.Name.Selector, func(e *colly.HTMLElement) {
			name := currSelector.Name.Parse(e)[0]

			if strings.HasSuffix(name, "") { // HQ icon; 3-byte character
				currRef.HQ = true
				name = name[0 : len(name)-3]
			}

			item := s.dataProvider.Item(name)
			if item != nil {
				currRef.NamedEntity = item
			}
		})

		materiaSelectors := []*selectors.SelectorInfo{
			&currSelector.Materia1,
			&currSelector.Materia2,
			&currSelector.Materia3,
			&currSelector.Materia4,
			&currSelector.Materia5,
		}
		for _, materiaSelector := range materiaSelectors {
			materiaCallback := func(e *colly.HTMLElement) {
				name := materiaSelector.ParseInnerHTML(e)[0]
				item := s.dataProvider.Item(name)
				if item != nil {
					currRef.Materia = append(currRef.Materia, item.ID)
				}
			}
			c.OnHTML(materiaSelector.Selector, materiaCallback)
		}

		c.OnHTML(currSelector.MirageName.Selector, func(e *colly.HTMLElement) {
			name := currSelector.MirageName.Parse(e)[0]
			item := s.dataProvider.Item(name)
			if item != nil {
				currRef.Mirage = item.ID
			}
		})
	}

	c.OnHTML(partSelectors.SoulCrystal.Name.Selector, func(e *colly.HTMLElement) {
		name := partSelectors.SoulCrystal.Name.Parse(e)[0]
		item := s.dataProvider.Item(name)
		if item != nil {
			partRefs.SoulCrystal.NamedEntity = item
		}
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
