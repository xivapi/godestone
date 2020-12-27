package collectors

import (
	"strconv"
	"time"

	lookups "github.com/karashiiro/godestone/table-lookups"

	"github.com/gocolly/colly/v2"
	"github.com/karashiiro/godestone/models"
	"github.com/karashiiro/godestone/pack/exports"
	"github.com/karashiiro/godestone/selectors"
)

// BuildFreeCompanySearchCollector builds the collector used for processing the page.
func BuildFreeCompanySearchCollector(
	meta *models.Meta,
	searchSelectors *selectors.SearchSelectors,
	grandCompanyTable *exports.GrandCompanyTable,
	output chan *models.FreeCompanySearchResult,
) *colly.Collector {
	c := colly.NewCollector(
		colly.MaxDepth(21),
		colly.UserAgent(meta.UserAgentDesktop),
		colly.IgnoreRobotsTxt(),
	)
	dur, _ := time.ParseDuration("60s")
	c.SetRequestTimeout(dur)

	fcSearchSelectors := searchSelectors.FreeCompany
	entrySelectors := fcSearchSelectors.Entry

	c.OnHTML(fcSearchSelectors.Root.Selector, func(container *colly.HTMLElement) {
		nextURI := fcSearchSelectors.ListNextButton.ParseThroughChildren(container)[0]

		container.ForEach(entrySelectors.Root.Selector, func(i int, e *colly.HTMLElement) {
			nextFC := models.FreeCompanySearchResult{
				Active: models.FreeCompanyActiveState(entrySelectors.Active.ParseThroughChildren(e)[0]),
				Name:   entrySelectors.Name.ParseThroughChildren(e)[0],
				ID:     entrySelectors.ID.ParseThroughChildren(e)[0],
				CrestLayers: &models.CrestLayers{
					Bottom: entrySelectors.CrestLayers.Bottom.ParseThroughChildren(e)[0],
					Middle: entrySelectors.CrestLayers.Middle.ParseThroughChildren(e)[0],
					Top:    entrySelectors.CrestLayers.Top.ParseThroughChildren(e)[0],
				},
				Estate:      entrySelectors.EstateBuilt.ParseThroughChildren(e)[0],
				Recruitment: models.FreeCompanyRecruitingState(entrySelectors.RecruitmentOpen.ParseThroughChildren(e)[0]),
			}

			gcName := entrySelectors.GrandCompany.ParseThroughChildren(e)[0]
			gc := lookups.GrandCompanyTableLookup(grandCompanyTable, gcName)

			nGCs := grandCompanyTable.GrandCompaniesLength()
			for i := 0; i < nGCs; i++ {
				nextFC.GrandCompany = &models.NamedEntity{
					ID:   gc.Id(),
					Name: gcName,

					NameEN: string(gc.NameEn()),
					NameJA: string(gc.NameJa()),
					NameDE: string(gc.NameDe()),
					NameFR: string(gc.NameFr()),
				}
			}

			server := entrySelectors.Server.ParseThroughChildren(e)
			nextFC.World = server[0]
			nextFC.DC = server[1]

			datetimeSecondsStr := entrySelectors.Formed.Parse(e)[0]
			datetimeSeconds, err := strconv.ParseInt(datetimeSecondsStr, 10, 32)
			if err == nil {
				nextFC.Formed = time.Unix(0, datetimeSeconds*1000*int64(time.Millisecond))
			}

			activeMembersStr := entrySelectors.ActiveMembers.ParseThroughChildren(e)[0]
			activeMembers, err := strconv.ParseUint(activeMembersStr, 10, 32)
			if err == nil {
				nextFC.ActiveMembers = uint32(activeMembers)
			}

			output <- &nextFC
		})

		if nextURI != "javascript:void(0);" && nextURI != "" /* "Your search yielded no results." */ {
			err := container.Request.Visit(nextURI)
			if err != nil {
				output <- &models.FreeCompanySearchResult{
					Error: err,
				}
			}
		}
	})

	return c
}
