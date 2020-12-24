package collectors

import (
	"strconv"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/karashiiro/godestone/models"
	"github.com/karashiiro/godestone/selectors"
)

// BuildLinkshellSearchCollector builds the collector used for processing the page.
func BuildLinkshellSearchCollector(meta *models.Meta, searchSelectors *selectors.SearchSelectors, output chan *models.LinkshellSearchResult) *colly.Collector {
	c := colly.NewCollector(
		colly.MaxDepth(21),
		colly.UserAgent(meta.UserAgentDesktop),
		colly.IgnoreRobotsTxt(),
	)
	dur, _ := time.ParseDuration("60s")
	c.SetRequestTimeout(dur)

	lsSearchSelectors := searchSelectors.Linkshell
	entrySelectors := lsSearchSelectors.Entry

	c.OnHTML(lsSearchSelectors.Root.Selector, func(container *colly.HTMLElement) {
		nextURI := lsSearchSelectors.ListNextButton.ParseThroughChildren(container)[0]

		container.ForEach(entrySelectors.Root.Selector, func(i int, e *colly.HTMLElement) {
			nextLinkshell := models.LinkshellSearchResult{
				Name: entrySelectors.Name.ParseThroughChildren(e)[0],
				ID:   entrySelectors.ID.ParseThroughChildren(e)[0],
			}

			server := entrySelectors.Server.ParseThroughChildren(e)
			nextLinkshell.World = server[0]
			nextLinkshell.DC = server[1]

			activeMembersStr := entrySelectors.ActiveMembers.ParseThroughChildren(e)[0]
			activeMembers, err := strconv.ParseUint(activeMembersStr, 10, 32)
			if err == nil {
				nextLinkshell.ActiveMembers = uint32(activeMembers)
			}

			output <- &nextLinkshell
		})

		if nextURI != "javascript:void(0);" {
			err := container.Request.Visit(nextURI)
			if err != nil {
				output <- &models.LinkshellSearchResult{
					Error: err,
				}
			}
		}
	})

	return c
}
