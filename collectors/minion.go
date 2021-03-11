package collectors

import (
	"github.com/gocolly/colly/v2"
	"github.com/karashiiro/godestone/models"
	"github.com/karashiiro/godestone/pack/exports"
	"github.com/karashiiro/godestone/selectors"
	lookups "github.com/karashiiro/godestone/table-lookups"
)

// BuildMinionCollector builds the collector used for processing the page.
func BuildMinionCollector(
	meta *models.Meta,
	profSelectors *selectors.ProfileSelectors,
	minionTable *exports.MinionTable,
	lang string,
	output chan *models.Minion,
) *colly.Collector {
	c := colly.NewCollector(
		colly.UserAgent(meta.UserAgentMobile),
		colly.IgnoreRobotsTxt(),
		colly.Async(),
	)

	minionSelectors := profSelectors.Minion

	c.OnHTML(minionSelectors.Minions.Root.Selector, func(e *colly.HTMLElement) {
		name := minionSelectors.Minions.Name.ParseThroughChildren(e)[0]
		icon := minionSelectors.Minions.Icon.ParseThroughChildren(e)[0]

		m := lookups.MinionTableLookup(minionTable, name, lang)
		if m == nil {
			output <- &models.Minion{
				ID:   0,
				Name: name,
				Icon: icon,

				NameEN: "",
				NameDE: "",
				NameFR: "",
				NameJA: "",
			}
		} else {
			output <- &models.Minion{
				ID:   m.Id(),
				Name: name,
				Icon: icon,

				NameEN: string(m.NameEn()),
				NameDE: string(m.NameDe()),
				NameFR: string(m.NameFr()),
				NameJA: string(m.NameJa()),
			}
		}
	})

	return c
}
