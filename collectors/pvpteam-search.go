package collectors

import (
	"strconv"

	"github.com/gocolly/colly/v2"
	"github.com/karashiiro/godestone/models"
	"github.com/karashiiro/godestone/selectors"
)

// BuildPVPTeamSearchCollector builds the collector used for processing the page.
func BuildPVPTeamSearchCollector(
	meta *models.Meta,
	searchSelectors *selectors.SearchSelectors,
	pageInfo *models.PageInfo,
	output chan *models.PVPTeamSearchResult,
) *colly.Collector {
	c := colly.NewCollector(
		colly.MaxDepth(2),
		colly.UserAgent(meta.UserAgentDesktop),
		colly.IgnoreRobotsTxt(),
		colly.AllowURLRevisit(),
	)

	pvpTeamSearchSelectors := searchSelectors.PVPTeam
	entrySelectors := pvpTeamSearchSelectors.Entry

	c.OnHTML(pvpTeamSearchSelectors.Root.Selector, func(container *colly.HTMLElement) {
		nextURI := pvpTeamSearchSelectors.ListNextButton.ParseThroughChildren(container)[0]

		pi := pvpTeamSearchSelectors.PageInfo.ParseThroughChildren(container)
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
			nextTeam := models.PVPTeamSearchResult{
				Name: entrySelectors.Name.ParseThroughChildren(e)[0],
				ID:   entrySelectors.ID.ParseThroughChildren(e)[0],
				DC:   entrySelectors.DC.ParseThroughChildren(e)[0],
				CrestLayers: &models.CrestLayers{
					Bottom: entrySelectors.CrestLayers.Bottom.ParseThroughChildren(e)[0],
					Middle: entrySelectors.CrestLayers.Middle.ParseThroughChildren(e)[0],
					Top:    entrySelectors.CrestLayers.Top.ParseThroughChildren(e)[0],
				},
			}

			output <- &nextTeam
		})

		revisited := false
		if !revisited && nextURI == "" {
			revisited = true
			err := c.Visit(container.Request.URL.String())
			if err != nil {
				output <- &models.PVPTeamSearchResult{
					Error: err,
				}
			}
		}
	})

	return c
}
