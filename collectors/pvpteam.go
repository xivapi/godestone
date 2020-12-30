package collectors

import (
	"strconv"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/karashiiro/godestone/data/gcrank"
	"github.com/karashiiro/godestone/models"
	"github.com/karashiiro/godestone/selectors"
)

// BuildPVPTeamCollector builds the collector used for processing the page.
func BuildPVPTeamCollector(meta *models.Meta, pvpTeamSelectors *selectors.PVPTeamSelectors, pvpTeam *models.PVPTeam) *colly.Collector {
	c := colly.NewCollector(
		colly.UserAgent(meta.UserAgentDesktop),
		colly.IgnoreRobotsTxt(),
		colly.Async(),
	)

	basicSelectors := pvpTeamSelectors.Basic
	c.OnHTML(basicSelectors.Name.Selector, func(e *colly.HTMLElement) {
		pvpTeam.Name = basicSelectors.Name.Parse(e)[0]
	})

	c.OnHTML(basicSelectors.DC.Selector, func(e *colly.HTMLElement) {
		pvpTeam.DC = basicSelectors.DC.Parse(e)[0]
	})

	c.OnHTML(basicSelectors.Formed.Selector, func(e *colly.HTMLElement) {
		datetimeSecondsStr := basicSelectors.Formed.Parse(e)[0]
		datetimeSeconds, err := strconv.ParseInt(datetimeSecondsStr, 10, 32)
		if err == nil {
			pvpTeam.Formed = time.Unix(0, datetimeSeconds*1000*int64(time.Millisecond))
		}
	})

	pvpTeam.CrestLayers = &models.CrestLayers{}
	c.OnHTML(basicSelectors.CrestLayers.Bottom.Selector, func(e *colly.HTMLElement) {
		pvpTeam.CrestLayers.Bottom = basicSelectors.CrestLayers.Bottom.Parse(e)[0]
	})

	c.OnHTML(basicSelectors.CrestLayers.Middle.Selector, func(e *colly.HTMLElement) {
		pvpTeam.CrestLayers.Middle = basicSelectors.CrestLayers.Middle.Parse(e)[0]
	})

	c.OnHTML(basicSelectors.CrestLayers.Top.Selector, func(e *colly.HTMLElement) {
		pvpTeam.CrestLayers.Top = basicSelectors.CrestLayers.Top.Parse(e)[0]
	})

	pvpTeam.Members = []*models.PVPTeamMember{}
	membersSelectors := pvpTeamSelectors.Members
	c.OnHTML(membersSelectors.Root.Selector, func(e1 *colly.HTMLElement) {
		e1.ForEach(membersSelectors.Entry.Root.Selector, func(i int, e2 *colly.HTMLElement) {
			member := &models.PVPTeamMember{
				Avatar:   membersSelectors.Entry.Avatar.ParseThroughChildren(e2)[0],
				Name:     membersSelectors.Entry.Name.ParseThroughChildren(e2)[0],
				Rank:     gcrank.Parse(membersSelectors.Entry.Rank.ParseThroughChildren(e2)[0]),
				RankIcon: membersSelectors.Entry.RankIcon.ParseThroughChildren(e2)[0],
			}

			worldDC := membersSelectors.Entry.Server.ParseThroughChildren(e2)
			member.World = worldDC[0]
			member.DC = worldDC[1]

			idStr := membersSelectors.Entry.ID.ParseThroughChildren(e2)[0]
			id, err := strconv.ParseUint(idStr, 10, 32)
			if err == nil {
				member.ID = uint32(id)
			}

			matchesStr := membersSelectors.Entry.Matches.ParseThroughChildren(e2)[0]
			matches, err := strconv.ParseUint(matchesStr, 10, 32)
			if err == nil {
				member.Matches = uint32(matches)
			}

			pvpTeam.Members = append(pvpTeam.Members, member)
		})
	})

	return c
}
