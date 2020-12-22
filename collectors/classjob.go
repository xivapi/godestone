package collectors

import (
	"github.com/gocolly/colly/v2"
	"github.com/karashiiro/godestone/models"
	"github.com/karashiiro/godestone/selectors"
)

// BuildClassJobCollector builds the collector used for processing the page.
func BuildClassJobCollector(meta *models.Meta, profSelectors *selectors.ProfileSelectors, charData *models.Character) *colly.Collector {
	c := colly.NewCollector()
	c.UserAgent = meta.UserAgentDesktop
	c.IgnoreRobotsTxt = true

	//classJobSelectors := profSelectors.ClassJob

	charData.ClassJobs = make([]*models.ClassJob, 0)

	return c
}
