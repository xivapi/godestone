package godestone

import (
	"log"

	"github.com/gocolly/colly/v2"
	"github.com/karashiiro/godestone/models"
)

func (s *Scraper) makeAchievementCollector(achievements *models.Achievements) *colly.Collector {
	c := colly.NewCollector()
	c.UserAgent = s.meta.UserAgentDesktop
	c.IgnoreRobotsTxt = true
	c.MaxDepth = 100 // Should be set to ceil(nAchievements / 50)

	achievementSelectors := s.profSelectors.Achievements

	achievements.List = make([]*models.AchievementInfo, 0)

	var nextButton *colly.HTMLElement
	c.OnHTML(achievementSelectors.ListNextButton.Selector, func(e *colly.HTMLElement) {
		nextButton = e
	})

	c.OnHTML(achievementSelectors.List.Selector, func(e1 *colly.HTMLElement) {
		e1.ForEach(achievementSelectors.ID.Selector, func(i int, e2 *colly.HTMLElement) {
			log.Println(e2.Attr("href"))
		})

		nextURI := nextButton.Attr("href")
		e1.Request.Visit(nextURI)
	})

	return c
}
