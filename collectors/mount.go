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
	c := colly.NewCollector()
	c.UserAgent = meta.UserAgentMobile
	c.IgnoreRobotsTxt = true

	mountSelectors := profSelectors.Mount

	c.OnHTML(mountSelectors.Mounts.Root.Selector, func(e *colly.HTMLElement) {
		name := mountSelectors.Mounts.Name.ParseThroughChildren(e)[0]
		icon := mountSelectors.Mounts.Icon.ParseThroughChildren(e)[0]

		m := lookups.MountTableLookup(mountTable, name)
		output <- &models.Mount{
			ID:   m.Id(),
			Name: name,
			Icon: icon,

			NameEN: string(m.NameEn()),
			NameDE: string(m.NameDe()),
			NameFR: string(m.NameFr()),
			NameJA: string(m.NameJa()),
		}
	})

	return c
}
