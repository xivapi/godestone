package collectors

import (
	"log"

	"github.com/gocolly/colly/v2"
	"github.com/karashiiro/godestone/models"
	"github.com/karashiiro/godestone/selectors"
)

// BuildMinionCollector builds the collector used for processing the page.
func BuildMinionCollector(meta *models.Meta, profSelectors *selectors.ProfileSelectors) *colly.Collector {
	c := colly.NewCollector()
	c.UserAgent = meta.UserAgentMobile
	c.IgnoreRobotsTxt = true

	minionSelectors := profSelectors.Minion

	c.OnHTML(minionSelectors.List.Selector, func(e1 *colly.HTMLElement) {
		e1.ForEach(minionSelectors.Name.Selector, func(i int, e2 *colly.HTMLElement) {
			log.Println(minionSelectors.Name.Parse(e2)[0])
		})
	})

	return c
}
