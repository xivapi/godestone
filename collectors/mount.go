package collectors

import (
	"strings"

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
		nameLower := strings.ToLower(name)

		icon := mountSelectors.Mounts.Icon.ParseThroughChildren(e)[0]

		nMounts := mountTable.MountsLength()
		for i := 0; i < nMounts; i++ {
			mount := exports.Mount{}
			mountTable.Mounts(&mount, i)

			nameEn := string(mount.NameEn())
			nameDe := string(mount.NameDe())
			nameFr := string(mount.NameFr())
			nameJa := string(mount.NameJa())

			nameEnLower := strings.ToLower(nameEn)
			nameDeLower := strings.ToLower(nameDe)
			nameFrLower := strings.ToLower(nameFr)
			nameJaLower := strings.ToLower(nameJa)

			if nameEnLower == nameLower || nameDeLower == nameLower || nameFrLower == nameLower || nameJaLower == nameLower {
				output <- &models.Mount{
					ID:   mount.Id(),
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
