package collectors

import (
	"strconv"

	"github.com/gocolly/colly/v2"
	"github.com/karashiiro/godestone/data/gcrank"
	"github.com/karashiiro/godestone/models"
	"github.com/karashiiro/godestone/selectors"
)

// BuildFreeCompanyMembersCollector builds the collector used for processing the page.
func BuildFreeCompanyMembersCollector(meta *models.Meta, fcSelectors *selectors.FreeCompanySelectors, output chan *models.FreeCompanyMember) *colly.Collector {
	c := colly.NewCollector(
		colly.UserAgent(meta.UserAgentDesktop),
		colly.MaxDepth(50),
		colly.IgnoreRobotsTxt(),
	)

	membersSelectors := fcSelectors.Members
	c.OnHTML(membersSelectors.Root.Selector, func(container *colly.HTMLElement) {
		nextURI := membersSelectors.ListNextButton.ParseThroughChildren(container)[0]

		container.ForEach(membersSelectors.Entry.Root.Selector, func(i int, e *colly.HTMLElement) {
			member := &models.FreeCompanyMember{
				Avatar:   membersSelectors.Entry.Avatar.ParseThroughChildren(e)[0],
				Name:     membersSelectors.Entry.Name.ParseThroughChildren(e)[0],
				Rank:     gcrank.Parse(membersSelectors.Entry.Rank.ParseThroughChildren(e)[0]),
				RankIcon: membersSelectors.Entry.RankIcon.ParseThroughChildren(e)[0],
			}

			worldDC := membersSelectors.Entry.Server.ParseThroughChildren(e)
			member.World = worldDC[0]
			member.DC = worldDC[1]

			idStr := membersSelectors.Entry.ID.ParseThroughChildren(e)[0]
			id, err := strconv.ParseUint(idStr, 10, 32)
			if err == nil {
				member.ID = uint32(id)
			}

			output <- member
		})

		if nextURI != "javascript:void(0);" {
			err := container.Request.Visit(nextURI)
			if err != nil {
				output <- &models.FreeCompanyMember{
					Error: err,
				}
			}
		}
	})

	return c
}
