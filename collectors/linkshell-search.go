package collectors

import (
	"strconv"

	"github.com/gocolly/colly/v2"
	"github.com/karashiiro/godestone/models"
	"github.com/karashiiro/godestone/selectors"
)

// BuildLinkshellSearchCollector builds the collector used for processing the page.
func BuildLinkshellSearchCollector(
	meta *models.Meta,
	searchSelectors *selectors.SearchSelectors,
	pageInfo *models.PageInfo,
	output chan *models.LinkshellSearchResult,
) *colly.Collector {
	c := colly.NewCollector(
		colly.MaxDepth(2),
		colly.UserAgent(meta.UserAgentDesktop),
		colly.IgnoreRobotsTxt(),
		colly.AllowURLRevisit(),
		colly.Async(),
	)

	lsSearchSelectors := searchSelectors.Linkshell
	entrySelectors := lsSearchSelectors.Entry

	c.OnHTML(lsSearchSelectors.Root.Selector, func(container *colly.HTMLElement) {
		nextURI := lsSearchSelectors.ListNextButton.ParseThroughChildren(container)[0]

		pi := lsSearchSelectors.PageInfo.ParseThroughChildren(container)
		if len(pi) > 1 {
			curPage, err := strconv.Atoi(pi[0])
			if err == nil {
				pageInfo.CurrentPage = curPage
			}
			totalPages, err := strconv.Atoi(pi[1])
			if err == nil {
				pageInfo.TotalPages = totalPages
			}
		}

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

		revisited := false
		if !revisited && nextURI == "" {
			revisited = true
			err := c.Visit(container.Request.URL.String())
			if err != nil {
				output <- &models.LinkshellSearchResult{
					Error: err,
				}
			}
		}
	})

	return c
}
