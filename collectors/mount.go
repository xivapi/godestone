package collectors

import (
	lookups "github.com/karashiiro/godestone/table-lookups"

	"github.com/gocolly/colly/v2"
	"github.com/karashiiro/godestone/models"
	"github.com/karashiiro/godestone/pack/exports"
	"github.com/karashiiro/godestone/selectors"
)

// BuildMountCollector builds the collector used for processing the page.
func BuildMountCollector(meta *models.Meta, profSelectors *selectors.ProfileSelectors, mountTable *exports.MountTable, output chan *models.Mount) *colly.Collector {
	c := colly.NewCollector(
		colly.UserAgent(meta.UserAgentMobile),
		colly.IgnoreRobotsTxt(),
		colly.Async(),
	)

	mountSelectors := profSelectors.Mount

	c.OnHTML(mountSelectors.Mounts.Root.Selector, func(e *colly.HTMLElement) {
		name := mountSelectors.Mounts.Name.ParseThroughChildren(e)[0]
		icon := mountSelectors.Mounts.Icon.ParseThroughChildren(e)[0]

		m := lookups.MountTableLookup(mountTable, name)
		if m == nil {
			output <- &models.Mount{
				ID:   0,
				Name: name,
				Icon: icon,

				NameEN: "",
				NameDE: "",
				NameFR: "",
				NameJA: "",
			}
		} else {
			output <- &models.Mount{
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
