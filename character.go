package godestone

import (
	"strconv"
	"strings"

	"github.com/gocolly/colly/v2"
	"github.com/karashiiro/godestone/data/baseparam"
	"github.com/karashiiro/godestone/data/gcrank"
	"github.com/karashiiro/godestone/data/gender"
	"github.com/karashiiro/godestone/selectors"
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

		mgc := s.grandCompanyTableLookup(gcName)
		gc := &NamedEntity{
			ID:   mgc.Id(),
			Name: gcName,

			NameEN: string(mgc.NameEn()),
			NameJA: string(mgc.NameJa()),
			NameDE: string(mgc.NameDe()),
			NameFR: string(mgc.NameFr()),
		}

		charData.GrandCompanyInfo = &GrandCompanyInfo{GrandCompany: gc, RankID: gcRank}
	})

	charData.GuardianDeity = &NamedEntity{}
	c.OnHTML(charSelectors.GuardianDeity.Name.Selector, func(e *colly.HTMLElement) {
		name := charSelectors.GuardianDeity.Name.Parse(e)[0]
		d := s.deityTableLookup(name)

		charData.GuardianDeity.ID = d.Id()
		charData.GuardianDeity.Name = name
		charData.GuardianDeity.NameEN = string(d.NameEn())
		charData.GuardianDeity.NameJA = string(d.NameJa())
		charData.GuardianDeity.NameDE = string(d.NameDe())
		charData.GuardianDeity.NameFR = string(d.NameFr())
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

		r := s.raceTableLookup(raceName)
		charData.Race = &GenderedEntity{
			ID:   r.Id(),
			Name: values[0],

			NameMasculineEN: string(r.NameMasculineEn()),
			NameFeminineEN:  string(r.NameFeminineEn()),
			NameMasculineJA: string(r.NameMasculineJa()),
			NameFeminineJA:  string(r.NameFeminineJa()),
			NameMasculineDE: string(r.NameMasculineDe()),
			NameFeminineDE:  string(r.NameFeminineDe()),
			NameMasculineFR: string(r.NameMasculineFr()),
			NameFeminineFR:  string(r.NameFeminineFr()),
		}

		t := s.tribeTableLookup(values[1])
		charData.Tribe = &GenderedEntity{
			ID:   t.Id(),
			Name: values[0],

			NameMasculineEN: string(t.NameMasculineEn()),
			NameFeminineEN:  string(t.NameFeminineEn()),
			NameMasculineJA: string(t.NameMasculineJa()),
			NameFeminineJA:  string(t.NameFeminineJa()),
			NameMasculineDE: string(t.NameMasculineDe()),
			NameFeminineDE:  string(t.NameFeminineDe()),
			NameMasculineFR: string(t.NameMasculineFr()),
			NameFeminineFR:  string(t.NameFeminineFr()),
		}

		charData.Gender = gender.Parse(values[2])
	})

	c.OnHTML(charSelectors.Server.Selector, func(e *colly.HTMLElement) {
		values := charSelectors.Server.Parse(e)

		charData.World = values[0]
		charData.DC = values[1]
	})

	c.OnHTML(charSelectors.Title.Selector, func(e *colly.HTMLElement) {
		titleText := charSelectors.Title.Parse(e)[0]
		t := s.titleTableLookup(titleText)

		charData.Title = &Title{
			GenderedEntity: &GenderedEntity{
				ID:   t.Id(),
				Name: titleText,

				NameMasculineEN: string(t.NameMasculineEn()),
				NameMasculineDE: string(t.NameMasculineDe()),
				NameMasculineFR: string(t.NameMasculineFr()),
				NameMasculineJA: string(t.NameMasculineJa()),
				NameFeminineEN:  string(t.NameFeminineEn()),
				NameFeminineDE:  string(t.NameFeminineDe()),
				NameFeminineFR:  string(t.NameFeminineFr()),
				NameFeminineJA:  string(t.NameFeminineJa()),
			},
			Prefix: t.IsPrefix(),
		}
	})

	charData.Town = &NamedEntity{}
	c.OnHTML(charSelectors.Town.Name.Selector, func(e *colly.HTMLElement) {
		name := charSelectors.Town.Name.Parse(e)[0]
		t := s.townTableLookup(name)

		charData.Town.ID = t.Id()
		charData.Town.Name = name
		charData.Town.NameEN = string(t.NameEn())
		charData.Town.NameJA = string(t.NameJa())
		charData.Town.NameDE = string(t.NameDe())
		charData.Town.NameFR = string(t.NameFr())
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
			item := s.itemTableLookup(name)
			if item != nil {
				currRef.Dye = item.Id()
			}
		})
		c.OnHTML(currSelector.Name.Selector, func(e *colly.HTMLElement) {
			name := currSelector.Name.Parse(e)[0]

			if strings.HasSuffix(name, "") { // HQ icon; 3-byte character
				currRef.HQ = true
				name = name[0 : len(name)-3]
			}

			item := s.itemTableLookup(name)
			if item != nil {
				currRef.Name = name
				currRef.ID = item.Id()
				currRef.NameEN = string(item.NameEn())
				currRef.NameJA = string(item.NameJa())
				currRef.NameDE = string(item.NameDe())
				currRef.NameFR = string(item.NameFr())
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
				item := s.itemTableLookup(name)
				if item != nil {
					currRef.Materia = append(currRef.Materia, item.Id())
				}
			}
			c.OnHTML(materiaSelector.Selector, materiaCallback)
		}

		c.OnHTML(currSelector.MirageName.Selector, func(e *colly.HTMLElement) {
			name := currSelector.MirageName.Parse(e)[0]
			item := s.itemTableLookup(name)
			if item != nil {
				currRef.Mirage = item.Id()
			}
		})
	}

	c.OnHTML(partSelectors.SoulCrystal.Name.Selector, func(e *colly.HTMLElement) {
		partRefs.SoulCrystal.Name = partSelectors.SoulCrystal.Name.Parse(e)[0]
		item := s.itemTableLookup(partRefs.SoulCrystal.Name)
		if item != nil {
			partRefs.SoulCrystal.ID = item.Id()
			partRefs.SoulCrystal.NameEN = string(item.NameEn())
			partRefs.SoulCrystal.NameJA = string(item.NameJa())
			partRefs.SoulCrystal.NameDE = string(item.NameDe())
			partRefs.SoulCrystal.NameFR = string(item.NameFr())
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