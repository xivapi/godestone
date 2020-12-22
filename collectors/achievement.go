package collectors

import (
	"strconv"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/karashiiro/godestone/models"
	"github.com/karashiiro/godestone/selectors"
)

// BuildAchievementCollector builds the collector used for processing the page.
func BuildAchievementCollector(meta *models.Meta, profSelectors *selectors.ProfileSelectors, output chan *models.AchievementInfo) *colly.Collector {
	c := colly.NewCollector()
	c.UserAgent = meta.UserAgentDesktop
	c.IgnoreRobotsTxt = true
	c.MaxDepth = 100 // Should be set to ceil(nAchievements / 50)

	achievementSelectors := profSelectors.Achievements

	nextURI := ""
	c.OnHTML(achievementSelectors.ListNextButton.Selector, func(e *colly.HTMLElement) {
		nextURI = achievementSelectors.ListNextButton.Parse(e)[0]
	})

	c.OnHTML(achievementSelectors.List.Selector, func(e1 *colly.HTMLElement) {
		e1.ForEach(achievementSelectors.Entry.Selector, func(i int, e2 *colly.HTMLElement) {
			nextAchievement := &models.AchievementInfo{}

			idStr := achievementSelectors.ID.ParseThroughChildren(e2)[0]
			id, err := strconv.ParseUint(idStr, 10, 32)
			if err == nil {
				nextAchievement.ID = uint32(id)
			}

			datetimeSecondsStr := achievementSelectors.Time.ParseThroughChildren(e2)[0]
			datetimeSeconds, err := strconv.ParseInt(datetimeSecondsStr, 10, 32)
			if err == nil {
				nextAchievement.Date = time.Unix(0, datetimeSeconds*1000*int64(time.Millisecond))
			}

			output <- nextAchievement
		})

		if nextURI != "javascript:void(0);" {
			err := e1.Request.Visit(nextURI)
			if err != nil {
				output <- &models.AchievementInfo{
					Error: err,
				}
			}
		}
	})

	return c
}
