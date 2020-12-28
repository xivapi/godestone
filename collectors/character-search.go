package collectors

import (
	"strconv"

	"github.com/gocolly/colly/v2"
	"github.com/karashiiro/godestone/data/gcrank"
	"github.com/karashiiro/godestone/models"
	"github.com/karashiiro/godestone/selectors"
)

// BuildCharacterSearchCollector builds the collector used for processing the page.
func BuildCharacterSearchCollector(
	meta *models.Meta,
	searchSelectors *selectors.SearchSelectors,
	pageInfo *models.PageInfo,
	output chan *models.CharacterSearchResult,
) *colly.Collector {
	c := colly.NewCollector(
		colly.MaxDepth(2),
		colly.UserAgent(meta.UserAgentDesktop),
		colly.IgnoreRobotsTxt(),
		colly.AllowURLRevisit(),
	)

	charSearchSelectors := searchSelectors.Character
	entrySelectors := charSearchSelectors.Entry

	c.OnHTML(charSearchSelectors.Root.Selector, func(container *colly.HTMLElement) {
		nextURI := charSearchSelectors.ListNextButton.ParseThroughChildren(container)[0]

		pi := charSearchSelectors.PageInfo.ParseThroughChildren(container)
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
			nextCharacter := models.CharacterSearchResult{
				Avatar:   entrySelectors.Avatar.ParseThroughChildren(e)[0],
				Name:     entrySelectors.Name.ParseThroughChildren(e)[0],
				Lang:     entrySelectors.Lang.ParseThroughChildren(e)[0],
				RankIcon: entrySelectors.RankIcon.ParseThroughChildren(e)[0],
			}

			idStr := entrySelectors.ID.ParseThroughChildren(e)[0]
			id, err := strconv.ParseUint(idStr, 10, 32)
			if err == nil {
				nextCharacter.ID = uint32(id)
			}

			gcRank := entrySelectors.Rank.ParseThroughChildren(e)[0]
			nextCharacter.Rank = gcrank.Parse(gcRank)

			worldDC := entrySelectors.Server.ParseThroughChildren(e)
			nextCharacter.World = worldDC[0]
			nextCharacter.DC = worldDC[1]

			output <- &nextCharacter
		})

		revisited := false
		if !revisited && nextURI == "" {
			revisited = true
			err := c.Visit(container.Request.URL.String())
			if err != nil {
				output <- &models.CharacterSearchResult{
					Error: err,
				}
			}
		}
	})

	return c
}
