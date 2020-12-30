package collectors

import (
	"strconv"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/karashiiro/godestone/data/reputation"
	"github.com/karashiiro/godestone/data/role"
	"github.com/karashiiro/godestone/models"
	"github.com/karashiiro/godestone/pack/exports"
	"github.com/karashiiro/godestone/selectors"
	lookups "github.com/karashiiro/godestone/table-lookups"
)

// BuildFreeCompanyCollector builds the collector used for processing the page.
func BuildFreeCompanyCollector(
	meta *models.Meta,
	fcSelectors *selectors.FreeCompanySelectors,
	grandCompanyTable *exports.GrandCompanyTable,
	fc *models.FreeCompany,
) *colly.Collector {
	c := colly.NewCollector(
		colly.UserAgent(meta.UserAgentDesktop),
		colly.IgnoreRobotsTxt(),
		colly.Async(),
	)

	basicSelectors := fcSelectors.Basic
	c.OnHTML(basicSelectors.ActiveState.Selector, func(e *colly.HTMLElement) {
		fc.Active = models.FreeCompanyActiveState(basicSelectors.ActiveState.Parse(e)[0])
	})
	c.OnHTML(basicSelectors.ActiveMemberCount.Selector, func(e *colly.HTMLElement) {
		membersStr := basicSelectors.ActiveMemberCount.Parse(e)[0]
		members, err := strconv.ParseUint(membersStr, 10, 32)
		if err == nil {
			fc.ActiveMemberCount = uint32(members)
		}
	})
	c.OnHTML(basicSelectors.Formed.Selector, func(e *colly.HTMLElement) {
		datetimeSecondsStr := basicSelectors.Formed.Parse(e)[0]
		datetimeSeconds, err := strconv.ParseInt(datetimeSecondsStr, 10, 32)
		if err == nil {
			fc.Formed = time.Unix(0, datetimeSeconds*1000*int64(time.Millisecond))
		}
	})
	c.OnHTML(basicSelectors.Name.Selector, func(e *colly.HTMLElement) {
		fc.Name = basicSelectors.Name.Parse(e)[0]
	})
	c.OnHTML(basicSelectors.Rank.Selector, func(e *colly.HTMLElement) {
		rankStr := basicSelectors.Rank.Parse(e)[0]
		rank, err := strconv.ParseUint(rankStr, 10, 8)
		if err == nil {
			fc.Rank = uint8(rank)
		}
	})
	c.OnHTML(basicSelectors.Recruitment.Selector, func(e *colly.HTMLElement) {
		fc.Recruitment = models.FreeCompanyRecruitingState(basicSelectors.Recruitment.Parse(e)[0])
	})
	c.OnHTML(basicSelectors.Server.Selector, func(e *colly.HTMLElement) {
		worldDC := basicSelectors.Server.Parse(e)

		fc.World = worldDC[0]
		fc.DC = worldDC[1]
	})
	c.OnHTML(basicSelectors.Slogan.Selector, func(e *colly.HTMLElement) {
		fc.Slogan = basicSelectors.Slogan.Parse(e)[0]
	})
	c.OnHTML(basicSelectors.Tag.Selector, func(e *colly.HTMLElement) {
		fc.Tag = basicSelectors.Tag.Parse(e)[0]
	})

	fc.CrestLayers = &models.CrestLayers{}
	c.OnHTML(basicSelectors.CrestLayers.Bottom.Selector, func(e *colly.HTMLElement) {
		fc.CrestLayers.Bottom = basicSelectors.CrestLayers.Bottom.Parse(e)[0]
	})
	c.OnHTML(basicSelectors.CrestLayers.Middle.Selector, func(e *colly.HTMLElement) {
		fc.CrestLayers.Middle = basicSelectors.CrestLayers.Middle.Parse(e)[0]
	})
	c.OnHTML(basicSelectors.CrestLayers.Top.Selector, func(e *colly.HTMLElement) {
		fc.CrestLayers.Top = basicSelectors.CrestLayers.Top.Parse(e)[0]
	})

	fc.Estate = &models.Estate{}
	c.OnHTML(basicSelectors.Estate.NoEstate.Selector, func(e *colly.HTMLElement) {
		fc.Estate = nil
	})
	c.OnHTML(basicSelectors.Estate.Greeting.Selector, func(e *colly.HTMLElement) {
		fc.Estate.Greeting = basicSelectors.Estate.Greeting.Parse(e)[0]
	})
	c.OnHTML(basicSelectors.Estate.Name.Selector, func(e *colly.HTMLElement) {
		fc.Estate.Name = basicSelectors.Estate.Name.Parse(e)[0]
	})
	c.OnHTML(basicSelectors.Estate.Plot.Selector, func(e *colly.HTMLElement) {
		fc.Estate.Plot = basicSelectors.Estate.Plot.Parse(e)[0]
	})

	fc.Ranking = &models.FreeCompanyRanking{}
	c.OnHTML(basicSelectors.Ranking.Monthly.Selector, func(e *colly.HTMLElement) {
		rankMonthStr := basicSelectors.Ranking.Monthly.Parse(e)[0]
		rankMonth, err := strconv.ParseUint(rankMonthStr, 10, 32)
		if err == nil {
			fc.Ranking.Monthly = uint32(rankMonth)
		}
	})
	c.OnHTML(basicSelectors.Ranking.Weekly.Selector, func(e *colly.HTMLElement) {
		rankWeekStr := basicSelectors.Ranking.Weekly.Parse(e)[0]
		rankWeek, err := strconv.ParseUint(rankWeekStr, 10, 32)
		if err == nil {
			fc.Ranking.Weekly = uint32(rankWeek)
		}
	})

	focusSelectors := fcSelectors.Focuses
	fc.Focus = []*models.FreeCompanyFocusInfo{}
	c.OnHTML(focusSelectors.NotSpecified.Selector, func(e *colly.HTMLElement) {
		fc.Focus = nil
	})

	focusSelectorsPtrs := []*selectors.FreeCompanyFocusSelectors{
		&focusSelectors.RolePlaying,
		&focusSelectors.Leveling,
		&focusSelectors.Casual,
		&focusSelectors.Hardcore,
		&focusSelectors.Dungeons,
		&focusSelectors.Guildhests,
		&focusSelectors.Trials,
		&focusSelectors.Raids,
		&focusSelectors.PVP,
	}
	for _, nextFocus := range focusSelectorsPtrs {
		curFocus := nextFocus

		info := &models.FreeCompanyFocusInfo{}
		c.OnHTML(curFocus.Icon.Selector, func(e *colly.HTMLElement) {
			info.Icon = curFocus.Icon.Parse(e)[0]
		})
		c.OnHTML(curFocus.Name.Selector, func(e *colly.HTMLElement) {
			info.Kind = models.FreeCompanyFocus(curFocus.Name.Parse(e)[0])
		})
		c.OnHTML(curFocus.Status.Selector, func(e *colly.HTMLElement) {
			// Dangerous; this can match if the regex is broken because the return value will be an empty string
			info.Status = curFocus.Status.Parse(e)[0] == ""
		})

		fc.Focus = append(fc.Focus, info)
	}

	seekingSelectors := fcSelectors.Seeking
	fc.Seeking = []*models.FreeCompanySeekingInfo{}
	c.OnHTML(seekingSelectors.NotSpecified.Selector, func(e *colly.HTMLElement) {
		fc.Seeking = nil
	})

	roleSelectorsPtrs := []*selectors.FreeCompanySeekingSelectors{
		&seekingSelectors.Tank,
		&seekingSelectors.Healer,
		&seekingSelectors.DPS,
		&seekingSelectors.Crafter,
		&seekingSelectors.Gatherer,
	}
	for _, nextRole := range roleSelectorsPtrs {
		curRole := nextRole

		info := &models.FreeCompanySeekingInfo{}
		c.OnHTML(curRole.Icon.Selector, func(e *colly.HTMLElement) {
			info.Icon = curRole.Icon.Parse(e)[0]
		})
		c.OnHTML(curRole.Name.Selector, func(e *colly.HTMLElement) {
			info.Kind = role.Parse(curRole.Name.Parse(e)[0])
		})
		c.OnHTML(curRole.Status.Selector, func(e *colly.HTMLElement) {
			// Dangerous; this can match if the regex is broken because the return value will be an empty string
			info.Status = curRole.Status.Parse(e)[0] == ""
		})

		fc.Seeking = append(fc.Seeking, info)
	}

	repSelectors := fcSelectors.Reputation
	repSelectorsPtrs := []*selectors.FreeCompanyAlignmentSelectors{
		&repSelectors.Maelstrom,
		&repSelectors.Adders,
		&repSelectors.Flames,
	}
	fc.Reputation = []*models.FreeCompanyReputation{}
	for _, nextRep := range repSelectorsPtrs {
		curRep := nextRep

		rep := &models.FreeCompanyReputation{}
		c.OnHTML(curRep.Name.Selector, func(e *colly.HTMLElement) {
			gcName := curRep.Name.Parse(e)[0]
			gc := lookups.GrandCompanyTableLookup(grandCompanyTable, gcName)

			rep.GrandCompany = &models.NamedEntity{
				ID:   gc.Id(),
				Name: gcName,

				NameEN: string(gc.NameEn()),
				NameJA: string(gc.NameJa()),
				NameDE: string(gc.NameDe()),
				NameFR: string(gc.NameFr()),
			}
		})
		c.OnHTML(curRep.Progress.Selector, func(e *colly.HTMLElement) {
			progressStr := curRep.Progress.Parse(e)[0]
			progress, err := strconv.ParseUint(progressStr, 10, 8)
			if err == nil {
				rep.Progress = uint8(progress)
			}
		})
		c.OnHTML(curRep.Rank.Selector, func(e *colly.HTMLElement) {
			rep.Rank = reputation.Parse(curRep.Rank.Parse(e)[0])
		})

		fc.Reputation = append(fc.Reputation, rep)
	}

	return c
}
