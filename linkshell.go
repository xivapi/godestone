package godestone

import (
	"strconv"

	"github.com/gocolly/colly/v2"
	"github.com/xivapi/godestone/data/gcrank"
)

func (s *Scraper) buildLinkshellCollector(ls *Linkshell) *colly.Collector {
	c := colly.NewCollector(
		colly.UserAgent(s.meta.UserAgentDesktop),
		colly.IgnoreRobotsTxt(),
		colly.Async(),
	)

	basicSelectors := s.getLinkshellSelectors().Basic
	c.OnHTML(basicSelectors.Name.Selector, func(e *colly.HTMLElement) {
		ls.Name = basicSelectors.Name.Parse(e)[0]
	})

	ls.Members = []*LinkshellMember{}
	membersSelectors := s.getLinkshellSelectors().Members
	c.OnHTML(membersSelectors.Root.Selector, func(container *colly.HTMLElement) {
		nextURI := membersSelectors.ListNextButton.ParseThroughChildren(container)[0]

		container.ForEach(membersSelectors.Entry.Root.Selector, func(i int, e *colly.HTMLElement) {
			member := &LinkshellMember{
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

func (s *Scraper) buildCWLSCollector(cwls *CWLS) *colly.Collector {
	c := colly.NewCollector(
		colly.UserAgent(s.meta.UserAgentDesktop),
		colly.IgnoreRobotsTxt(),
		colly.Async(),
	)

	basicSelectors := s.getCWLSSelectors().Basic
	c.OnHTML(basicSelectors.Name.Selector, func(e *colly.HTMLElement) {
		cwls.Name = basicSelectors.Name.Parse(e)[0]
	})

	c.OnHTML(basicSelectors.DC.Selector, func(e *colly.HTMLElement) {
		cwls.DC = basicSelectors.DC.Parse(e)[0]
	})

	cwls.Members = []*CWLSMember{}
	membersSelectors := s.getCWLSSelectors().Members
	c.OnHTML(membersSelectors.Root.Selector, func(container *colly.HTMLElement) {
		nextURI := membersSelectors.ListNextButton.ParseThroughChildren(container)[0]

		container.ForEach(membersSelectors.Entry.Root.Selector, func(i int, e *colly.HTMLElement) {
			member := &CWLSMember{
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
