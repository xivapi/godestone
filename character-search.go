package godestone

import (
	"log"

	"github.com/gocolly/colly/v2"
)

func (s *Scraper) makeCharacterSearchCollector(name string, world string) *colly.Collector {
	c := colly.NewCollector()
	c.UserAgent = s.meta.UserAgentDesktop
	c.IgnoreRobotsTxt = true
	c.MaxDepth = 20

	charSearchSelectors := s.searchSelectors.Character

	c.OnHTML(charSearchSelectors.Entries.Selector, func(e *colly.HTMLElement) {
		name := e.DOM.ChildrenFiltered(charSearchSelectors.Entry.Name.Selector)
		log.Println(name.Text())
	})

	return c
}
