package godestone

import (
	"github.com/gocolly/colly/v2"
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

		m := s.minionTableLookup(name)
		if m == nil {
			output <- &Minion{
				ID:   0,
				Name: name,
				Icon: icon,

				NameEN: "",
				NameDE: "",
				NameFR: "",
				NameJA: "",
			}
		} else {
			output <- &Minion{
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

		m := s.mountTableLookup(name)
		if m == nil {
			output <- &Mount{
				ID:   0,
				Name: name,
				Icon: icon,

				NameEN: "",
				NameDE: "",
				NameFR: "",
				NameJA: "",
			}
		} else {
			output <- &Mount{
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
