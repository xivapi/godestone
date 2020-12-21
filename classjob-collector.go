package godestone

import (
	"github.com/gocolly/colly/v2"
	"github.com/karashiiro/godestone/models"
)

func (s *Scraper) makeClassJobCollector(charData *models.Character) *colly.Collector {
	c := colly.NewCollector()
	c.UserAgent = s.meta.UserAgentDesktop
	c.IgnoreRobotsTxt = true

	//classJobSelectors := s.profSelectors.ClassJob

	charData.ClassJobs = make([]*models.ClassJob, 0)

	return c
}
