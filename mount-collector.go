package godestone

import (
	"log"

	"github.com/gocolly/colly/v2"
)

func (s *Scraper) makeMountCollector() *colly.Collector {
	c := colly.NewCollector()
	c.UserAgent = s.meta.UserAgentMobile
	c.IgnoreRobotsTxt = true

	mountSelectors := s.profileSelectors.Mount

	c.OnHTML(mountSelectors.List.Selector, func(e1 *colly.HTMLElement) {
		e1.ForEach(mountSelectors.Name.Selector, func(i int, e2 *colly.HTMLElement) {
			log.Println(mountSelectors.Name.Parse(e2)[0])
		})
	})

	return c
}
