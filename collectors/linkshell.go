package collectors

import (
	"strconv"

	"github.com/gocolly/colly/v2"
	"github.com/karashiiro/godestone/data/gcrank"
	"github.com/karashiiro/godestone/models"
	"github.com/karashiiro/godestone/selectors"
)

// BuildLinkshellCollector builds the collector used for processing the page.
func BuildLinkshellCollector(meta *models.Meta, lsSelectors *selectors.LinkshellSelectors, ls *models.Linkshell) *colly.Collector {
	c := colly.NewCollector(
		colly.UserAgent(meta.UserAgentDesktop),
		colly.IgnoreRobotsTxt(),
		colly.Async(),
	)

	basicSelectors := lsSelectors.Basic
	c.OnHTML(basicSelectors.Name.Selector, func(e *colly.HTMLElement) {
		ls.Name = basicSelectors.Name.Parse(e)[0]
	})

	ls.Members = []*models.LinkshellMember{}
	membersSelectors := lsSelectors.Members
	c.OnHTML(membersSelectors.Root.Selector, func(container *colly.HTMLElement) {
		nextURI := membersSelectors.ListNextButton.ParseThroughChildren(container)[0]

		container.ForEach(membersSelectors.Entry.Root.Selector, func(i int, e *colly.HTMLElement) {
			member := &models.LinkshellMember{
				Avatar:            membersSelectors.Entry.Avatar.ParseThroughChildren(e)[0],
				Name:              membersSelectors.Entry.Name.ParseThroughChildren(e)[0],
				Rank:              gcrank.Parse(membersSelectors.Entry.Rank.ParseThroughChildren(e)[0]),
				RankIcon:          membersSelectors.Entry.RankIcon.ParseThroughChildren(e)[0],
				LinkshellRank:     membersSelectors.Entry.LinkshellRank.ParseThroughChildren(e)[0],
				LinkshellRankIcon: membersSelectors.Entry.LinkshellRankIcon.ParseThroughChildren(e)[0],
			}

			worldDC := membersSelectors.Entry.Server.ParseThroughChildren(e)
			member.World = worldDC[0]
			member.DC = worldDC[1]

			idStr := membersSelectors.Entry.ID.ParseThroughChildren(e)[0]
			id, err := strconv.ParseUint(idStr, 10, 32)
			if err == nil {
				member.ID = uint32(id)
			}

			ls.Members = append(ls.Members, member)
		})

		if nextURI != "javascript:void(0);" {
			container.Request.Visit(nextURI)
		}
	})

	return c
}
