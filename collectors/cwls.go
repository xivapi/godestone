package collectors

import (
	"strconv"

	"github.com/gocolly/colly/v2"
	"github.com/karashiiro/godestone/data/gcrank"
	"github.com/karashiiro/godestone/models"
	"github.com/karashiiro/godestone/selectors"
)

// BuildCWLSCollector builds the collector used for processing the page.
func BuildCWLSCollector(meta *models.Meta, cwlsSelectors *selectors.CWLSSelectors, cwls *models.CWLS) *colly.Collector {
	c := colly.NewCollector(
		colly.UserAgent(meta.UserAgentDesktop),
		colly.IgnoreRobotsTxt(),
	)

	basicSelectors := cwlsSelectors.Basic
	c.OnHTML(basicSelectors.Name.Selector, func(e *colly.HTMLElement) {
		cwls.Name = basicSelectors.Name.Parse(e)[0]
	})

	c.OnHTML(basicSelectors.DC.Selector, func(e *colly.HTMLElement) {
		cwls.DC = basicSelectors.DC.Parse(e)[0]
	})

	cwls.Members = []*models.CWLSMember{}
	membersSelectors := cwlsSelectors.Members
	c.OnHTML(membersSelectors.Root.Selector, func(container *colly.HTMLElement) {
		nextURI := membersSelectors.ListNextButton.ParseThroughChildren(container)[0]

		container.ForEach(membersSelectors.Entry.Root.Selector, func(i int, e *colly.HTMLElement) {
			member := &models.CWLSMember{
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

			cwls.Members = append(cwls.Members, member)
		})

		if nextURI != "javascript:void(0);" {
			container.Request.Visit(nextURI)
		}
	})

	return c
}
