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
		colly.MaxDepth(20),
		colly.UserAgent(s.meta.UserAgentDesktop),
		colly.IgnoreRobotsTxt(),
	)
	dur, _ := time.ParseDuration("60s")
	c.SetRequestTimeout(dur)

	charSearchSelectors := s.searchSelectors.Character
	entrySelectors := charSearchSelectors.Entry

	c.OnHTML(charSearchSelectors.EntriesContainer.Selector, func(container *colly.HTMLElement) {
		nextURI, uriExists := container.DOM.Find(charSearchSelectors.ListNextButton.Selector).Attr("href")

		container.ForEach(entrySelectors.Root.Selector, func(i int, e *colly.HTMLElement) {
			nextCharacter := models.CharacterSearchResult{
				Name:   entrySelectors.Name.ParseThroughChildren(e)[0],
				Lang:   entrySelectors.Lang.ParseThroughChildren(e)[0],
				Server: entrySelectors.Server.ParseThroughChildren(e)[0],
			}

			idStr := entrySelectors.ID.ParseThroughChildren(e)[0]
			id, err := strconv.ParseUint(idStr, 10, 32)
			if err == nil {
				nextCharacter.ID = uint32(id)
			}

			gcRank := entrySelectors.Rank.ParseThroughChildren(e)[0]
			nextCharacter.Rank = gcrank.Parse(gcRank)

			rankIcon := entrySelectors.RankIcon.ParseThroughChildren(e)[0]
			nextCharacter.RankIcon = rankIcon

			output <- &nextCharacter
		})

		if uriExists {
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
