package godestone

import (
	"github.com/gocolly/colly/v2"
	"github.com/xivapi/godestone/v2/provider/models"
)

func (s *Scraper) buildMinionCollector(output chan *Minion) *colly.Collector {
	c := colly.NewCollector(
		colly.UserAgent(s.meta.UserAgentMobile),
		colly.IgnoreRobotsTxt(),
		colly.Async(),
	)

	minionSelectors := s.getProfileSelectors().Minion

	c.OnHTML(minionSelectors.Minions.Root.Selector, func(e *colly.HTMLElement) {
		name := minionSelectors.Minions.Name.ParseThroughChildren(e)[0]
		icon := minionSelectors.Minions.Icon.ParseThroughChildren(e)[0]

		m, _ := s.dataProvider.Minion(name)
		if m == nil {
			output <- &Minion{
				IconedNamedEntity: &IconedNamedEntity{
					NamedEntity: &models.NamedEntity{
						ID:   0,
						Name: name,

						NameEN: "",
						NameDE: "",
						NameFR: "",
						NameJA: "",
					},
					Icon: icon,
				},
			}
		} else {
			output <- &Minion{
				IconedNamedEntity: &IconedNamedEntity{
					NamedEntity: m,
					Icon:        icon,
				},
			}
		}
	})

	return c
}

func (s *Scraper) buildMountCollector(output chan *Mount) *colly.Collector {
	c := colly.NewCollector(
		colly.UserAgent(s.meta.UserAgentMobile),
		colly.IgnoreRobotsTxt(),
		colly.Async(),
	)

	mountSelectors := s.getProfileSelectors().Mount

	c.OnHTML(mountSelectors.Mounts.Root.Selector, func(e *colly.HTMLElement) {
		name := mountSelectors.Mounts.Name.ParseThroughChildren(e)[0]
		icon := mountSelectors.Mounts.Icon.ParseThroughChildren(e)[0]

		m, _ := s.dataProvider.Mount(name)
		if m == nil {
			output <- &Mount{
				IconedNamedEntity: &IconedNamedEntity{
					NamedEntity: &models.NamedEntity{
						ID:   0,
						Name: name,

						NameEN: "",
						NameDE: "",
						NameFR: "",
						NameJA: "",
					},
					Icon: icon,
				},
			}
		} else {
			output <- &Mount{
				IconedNamedEntity: &IconedNamedEntity{
					NamedEntity: m,
					Icon:        icon,
				},
			}
		}
	})

	return c
}
