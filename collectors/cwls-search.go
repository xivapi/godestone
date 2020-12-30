package collectors

import (
	"strconv"

	"github.com/gocolly/colly/v2"
	"github.com/karashiiro/godestone/models"
	"github.com/karashiiro/godestone/selectors"
)

// BuildCWLSSearchCollector builds the collector used for processing the page.
func BuildCWLSSearchCollector(
	meta *models.Meta,
	searchSelectors *selectors.SearchSelectors,
	pageInfo *models.PageInfo,
	output chan *models.CWLSSearchResult,
) *colly.Collector {
	c := colly.NewCollector(
		colly.MaxDepth(2),
		colly.UserAgent(meta.UserAgentDesktop),
		colly.IgnoreRobotsTxt(),
		colly.AllowURLRevisit(),
		colly.Async(),
	)

	cwlsSearchSelectors := searchSelectors.CWLS
	entrySelectors := cwlsSearchSelectors.Entry

	c.OnHTML(cwlsSearchSelectors.Root.Selector, func(container *colly.HTMLElement) {
		nextURI := cwlsSearchSelectors.ListNextButton.ParseThroughChildren(container)[0]

		pi := cwlsSearchSelectors.PageInfo.ParseThroughChildren(container)
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
			nextCWLS := models.CWLSSearchResult{
				Name: entrySelectors.Name.ParseThroughChildren(e)[0],
				ID:   entrySelectors.ID.ParseThroughChildren(e)[0],
				DC:   entrySelectors.DC.ParseThroughChildren(e)[0],
			}

			activeMembersStr := entrySelectors.ActiveMembers.ParseThroughChildren(e)[0]
			activeMembers, err := strconv.ParseUint(activeMembersStr, 10, 32)
			if err == nil {
				nextCWLS.ActiveMembers = uint32(activeMembers)
			}

			output <- &nextCWLS
		})

		revisited := false
		if !revisited && nextURI == "" {
			revisited = true
			err := c.Visit(container.Request.URL.String())
			if err != nil {
				output <- &models.CWLSSearchResult{
					Error: err,
				}
			}
		}
	})

	return c
}
