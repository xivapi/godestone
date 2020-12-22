package collectors

import (
	"log"

	"github.com/gocolly/colly/v2"
	"github.com/karashiiro/godestone/models"
	"github.com/karashiiro/godestone/selectors"
)

// BuildMountCollector builds the collector used for processing the page.
func BuildMountCollector(meta *models.Meta, profSelectors *selectors.ProfileSelectors) *colly.Collector {
	c := colly.NewCollector()
	c.UserAgent = meta.UserAgentMobile
	c.IgnoreRobotsTxt = true

	mountSelectors := profSelectors.Mount

	c.OnHTML(mountSelectors.List.Selector, func(e1 *colly.HTMLElement) {
		e1.ForEach(mountSelectors.Name.Selector, func(i int, e2 *colly.HTMLElement) {
			log.Println(mountSelectors.Name.Parse(e2)[0])
		})
	})

	return c
}
