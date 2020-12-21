package godestone

import (
	"strconv"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/karashiiro/godestone/data/gcrank"
	"github.com/karashiiro/godestone/models"
)

func (s *Scraper) makeCharacterSearchCollector(output chan *models.CharacterSearchResult) *colly.Collector {
	c := colly.NewCollector(
		colly.MaxDepth(21),
		colly.UserAgent(s.meta.UserAgentDesktop),
		colly.IgnoreRobotsTxt(),
	)
	dur, _ := time.ParseDuration("60s")
	c.SetRequestTimeout(dur)

	charSearchSelectors := s.searchSelectors.Character
	entrySelectors := charSearchSelectors.Entry

	c.OnHTML(charSearchSelectors.EntriesContainer.Selector, func(container *colly.HTMLElement) {
		nextURI := charSearchSelectors.ListNextButton.ParseThroughChildren(container)[0]

		container.ForEach(entrySelectors.Root.Selector, func(i int, e *colly.HTMLElement) {
			nextCharacter := models.CharacterSearchResult{
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
			nextCharacter.Server = worldDC[0]
			nextCharacter.DC = worldDC[1]

			output <- &nextCharacter
		})

		if nextURI != "" {
			err := container.Request.Visit(nextURI)
			if err != nil {
				output <- &models.CharacterSearchResult{
					Error: err,
				}
			}
		}
	})

	return c
}
