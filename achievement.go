package godestone

import (
	"strconv"
	"time"

	"github.com/gocolly/colly/v2"
)

func (s *Scraper) buildAchievementCollector(output chan *AchievementInfo) *colly.Collector {
	c := colly.NewCollector(
		colly.UserAgent(s.meta.UserAgentDesktop),
		colly.IgnoreRobotsTxt(),
		colly.MaxDepth(100), // Should be set to ceil(nAchievements / 50) + 1
		colly.Async(),
	)

	achievementSelectors := s.getProfileSelectors().Achievements

	allAchievementInfo := &AllAchievementInfo{}
	c.OnHTML(achievementSelectors.TotalAchievements.Selector, func(e *colly.HTMLElement) {
		taStr := achievementSelectors.TotalAchievements.Parse(e)[0]
		ta, err := strconv.ParseUint(taStr, 10, 32)
		if err == nil {
			allAchievementInfo.TotalAchievements = uint32(ta)
		}
	})
	c.OnHTML(achievementSelectors.AchievementPoints.Selector, func(e *colly.HTMLElement) {
		apStr := achievementSelectors.AchievementPoints.Parse(e)[0]
		ap, err := strconv.ParseUint(apStr, 10, 32)
		if err == nil {
			allAchievementInfo.TotalAchievementPoints = uint32(ap)
		}
	})

	c.OnHTML(achievementSelectors.Root.Selector, func(e1 *colly.HTMLElement) {
		nextURI := achievementSelectors.ListNextButton.ParseThroughChildren(e1)[0]

		entrySelectors := achievementSelectors.Entry
		e1.ForEach(entrySelectors.Root.Selector, func(i int, e2 *colly.HTMLElement) {
			nameOptions := entrySelectors.Name.ParseThroughChildren(e2)
			name := nameOptions[0]
			if name == "" {
				name = nameOptions[1]
			}

			nextAchievement := &AchievementInfo{
				AllAchievementInfo: allAchievementInfo,
				Name:               name,
			}

			achievement := s.achievementTableLookup(name)
			if achievement != nil {
				nextAchievement.NameEN = string(achievement.NameEn())
				nextAchievement.NameJA = string(achievement.NameJa())
				nextAchievement.NameDE = string(achievement.NameDe())
				nextAchievement.NameFR = string(achievement.NameFr())
			}

			idStr := entrySelectors.ID.ParseThroughChildren(e2)[0]
			id, err := strconv.ParseUint(idStr, 10, 32)
			if err == nil {
				nextAchievement.ID = uint32(id)
			}

			datetimeSecondsStr := entrySelectors.Time.ParseThroughChildren(e2)[0]
			datetimeSeconds, err := strconv.ParseInt(datetimeSecondsStr, 10, 32)
			if err == nil {
				nextAchievement.Date = time.Unix(0, datetimeSeconds*1000*int64(time.Millisecond))
			}

			output <- nextAchievement
		})

		if nextURI != "javascript:void(0);" {
			err := e1.Request.Visit(nextURI)
			if err != nil {
				output <- &AchievementInfo{
					Error: err,
				}
			}
		}
	})

	return c
}
