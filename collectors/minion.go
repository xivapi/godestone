package collectors

import (
	"strings"

	"github.com/gocolly/colly/v2"
	"github.com/karashiiro/godestone/models"
	"github.com/karashiiro/godestone/pack/exports"
	"github.com/karashiiro/godestone/selectors"
)

// BuildMinionCollector builds the collector used for processing the page.
func BuildMinionCollector(meta *models.Meta, profSelectors *selectors.ProfileSelectors, minionTable *exports.MinionTable, output chan *models.Minion) *colly.Collector {
	c := colly.NewCollector()
	c.UserAgent = meta.UserAgentMobile
	c.IgnoreRobotsTxt = true

	minionSelectors := profSelectors.Minion

	c.OnHTML(minionSelectors.Minions.Root.Selector, func(e *colly.HTMLElement) {
		name := minionSelectors.Minions.Name.ParseThroughChildren(e)[0]
		nameLower := strings.ToLower(name)

		icon := minionSelectors.Minions.Icon.ParseThroughChildren(e)[0]

		nMinions := minionTable.MinionsLength()
		for i := 0; i < nMinions; i++ {
			minion := exports.Minion{}
			minionTable.Minions(&minion, i)

			nameEn := string(minion.NameEn())
			nameDe := string(minion.NameDe())
			nameFr := string(minion.NameFr())
			nameJa := string(minion.NameJa())

			nameEnLower := strings.ToLower(nameEn)
			nameDeLower := strings.ToLower(nameDe)
			nameFrLower := strings.ToLower(nameFr)
			nameJaLower := strings.ToLower(nameJa)

			if nameEnLower == nameLower || nameDeLower == nameLower || nameFrLower == nameLower || nameJaLower == nameLower {
				output <- &models.Minion{
					ID:   minion.Id(),
					Name: name,
					Icon: icon,

					NameEN: nameEn,
					NameDE: nameDe,
					NameFR: nameFr,
					NameJA: nameJa,
				}
			}
		}
	})

	return c
}
