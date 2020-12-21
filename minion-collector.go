package godestone

import (
	"log"

	"github.com/gocolly/colly/v2"
)

func (s *Scraper) makeMinionCollector() *colly.Collector {
	c := colly.NewCollector()
	c.UserAgent = s.meta.UserAgentMobile
	c.IgnoreRobotsTxt = true

	minionSelectors := s.profileSelectors.Minion

	c.OnHTML(minionSelectors.List.Selector, func(e1 *colly.HTMLElement) {
		e1.ForEach(minionSelectors.Name.Selector, func(i int, e2 *colly.HTMLElement) {
			log.Println(e2.Text)
		})
	})

	return c
}
