package collectors

import (
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/karashiiro/godestone/models"
	"github.com/karashiiro/godestone/pack/exports"
	"github.com/karashiiro/godestone/selectors"
)

// BuildAchievementCollector builds the collector used for processing the page.
func BuildAchievementCollector(meta *models.Meta, profSelectors *selectors.ProfileSelectors, achievementTable *exports.AchievementTable, output chan *models.AchievementInfo) *colly.Collector {
	c := colly.NewCollector(
		colly.UserAgent(meta.UserAgentDesktop),
		colly.IgnoreRobotsTxt(),
		colly.MaxDepth(100), // Should be set to ceil(nAchievements / 50) + 1
	)

	achievementSelectors := profSelectors.Achievements

	allAchievementInfo := &models.AllAchievementInfo{}
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
			nameLower := strings.ToLower(name)

			nextAchievement := &models.AchievementInfo{
				AllAchievementInfo: allAchievementInfo,
				Name:               name,
			}

			nAchievements := achievementTable.AchievementsLength()
			for i := 0; i < nAchievements; i++ {
				achievement := exports.Achievement{}
				achievementTable.Achievements(&achievement, i)

				nameEn := string(achievement.NameEn())
				nameDe := string(achievement.NameDe())
				nameFr := string(achievement.NameFr())
				nameJa := string(achievement.NameJa())

				nameEnLower := strings.ToLower(nameEn)
				nameDeLower := strings.ToLower(nameDe)
				nameFrLower := strings.ToLower(nameFr)
				nameJaLower := strings.ToLower(nameJa)

				if nameEnLower == nameLower || nameDeLower == nameLower || nameFrLower == nameLower || nameJaLower == nameLower {
					nextAchievement.NameEN = nameEn
					nextAchievement.NameJA = nameJa
					nextAchievement.NameDE = nameDe
					nextAchievement.NameFR = nameFr
				}
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
				output <- &models.AchievementInfo{
					Error: err,
				}
			}
		}
	})

	return c
}
