package godestone

import (
	"strconv"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/xivapi/godestone/v2/data/gcrank"
)

func (s *Scraper) buildCharacterSearchCollector(
	pageInfo *PageInfo,
	output chan *CharacterSearchResult,
) *colly.Collector {
	c := colly.NewCollector(
		colly.MaxDepth(2),
		colly.UserAgent(s.meta.UserAgentDesktop),
		colly.IgnoreRobotsTxt(),
		colly.AllowURLRevisit(),
		colly.Async(),
	)

	charSearchSelectors := s.getSearchSelectors().Character
	entrySelectors := charSearchSelectors.Entry

	c.OnHTML(charSearchSelectors.Root.Selector, func(container *colly.HTMLElement) {
		nextURI := charSearchSelectors.ListNextButton.ParseThroughChildren(container)[0]

		pi := charSearchSelectors.PageInfo.ParseThroughChildren(container)
		if len(pi) > 1 {
			curPage, err := strconv.Atoi(pi[0])
			if err == nil {
				pageInfo.CurrentPage = curPage
			}
			totalPages, err := strconv.Atoi(pi[1])
			if err == nil {
				pageInfo.TotalPages = totalPages
			}
		}

		container.ForEach(entrySelectors.Root.Selector, func(i int, e *colly.HTMLElement) {
			nextCharacter := CharacterSearchResult{
				Avatar:   entrySelectors.Avatar.ParseThroughChildren(e)[0],
				Name:     entrySelectors.Name.ParseThroughChildren(e)[0],
				Lang:     entrySelectors.Lang.ParseThroughChildren(e)[0],
				RankIcon: entrySelectors.RankIcon.ParseThroughChildren(e)[0],
			}

			idStr := entrySelectors.ID.ParseThroughChildren(e)[0]
			id, err := strconv.ParseUint(idStr, 10, 32)
			if err == nil {
				nextCharacter.ID = uint32(id)
			}

			gcRank := entrySelectors.Rank.ParseThroughChildren(e)[0]
			nextCharacter.Rank = gcrank.Parse(gcRank)

			worldDC := entrySelectors.Server.ParseThroughChildren(e)
			nextCharacter.World = worldDC[0]
			nextCharacter.DC = worldDC[1]

			output <- &nextCharacter
		})

		revisited := false
		if !revisited && nextURI == "" {
			revisited = true
			err := c.Visit(container.Request.URL.String())
			if err != nil {
				output <- &CharacterSearchResult{
					Error: err,
				}
			}
		}
	})

	return c
}

func (s *Scraper) buildCWLSSearchCollector(
	pageInfo *PageInfo,
	output chan *CWLSSearchResult,
) *colly.Collector {
	c := colly.NewCollector(
		colly.MaxDepth(2),
		colly.UserAgent(s.meta.UserAgentDesktop),
		colly.IgnoreRobotsTxt(),
		colly.AllowURLRevisit(),
		colly.Async(),
	)

	cwlsSearchSelectors := s.getSearchSelectors().CWLS
	entrySelectors := cwlsSearchSelectors.Entry

	c.OnHTML(cwlsSearchSelectors.Root.Selector, func(container *colly.HTMLElement) {
		nextURI := cwlsSearchSelectors.ListNextButton.ParseThroughChildren(container)[0]

		pi := cwlsSearchSelectors.PageInfo.ParseThroughChildren(container)
		if len(pi) > 1 {
			curPage, err := strconv.Atoi(pi[0])
			if err == nil {
				pageInfo.CurrentPage = curPage
			}
			totalPages, err := strconv.Atoi(pi[1])
			if err == nil {
				pageInfo.TotalPages = totalPages
			}
		}

		container.ForEach(entrySelectors.Root.Selector, func(i int, e *colly.HTMLElement) {
			nextCWLS := CWLSSearchResult{
				Name: entrySelectors.Name.ParseThroughChildren(e)[0],
				ID:   entrySelectors.ID.ParseThroughChildren(e)[0],
				DC:   entrySelectors.DC.ParseThroughChildren(e)[0],
			}

			activeMembersStr := entrySelectors.ActiveMembers.ParseThroughChildren(e)[0]
			activeMembers, err := strconv.ParseUint(activeMembersStr, 10, 32)
			if err == nil {
				nextCWLS.ActiveMembers = uint32(activeMembers)
			}

			output <- &nextCWLS
		})

		revisited := false
		if !revisited && nextURI == "" {
			revisited = true
			err := c.Visit(container.Request.URL.String())
			if err != nil {
				output <- &CWLSSearchResult{
					Error: err,
				}
			}
		}
	})

	return c
}

func (s *Scraper) buildFreeCompanySearchCollector(
	pageInfo *PageInfo,
	output chan *FreeCompanySearchResult,
) *colly.Collector {
	c := colly.NewCollector(
		colly.MaxDepth(2),
		colly.UserAgent(s.meta.UserAgentDesktop),
		colly.IgnoreRobotsTxt(),
		colly.AllowURLRevisit(),
		colly.Async(),
	)

	fcSearchSelectors := s.getSearchSelectors().FreeCompany
	entrySelectors := fcSearchSelectors.Entry

	c.OnHTML(fcSearchSelectors.Root.Selector, func(container *colly.HTMLElement) {
		nextURI := fcSearchSelectors.ListNextButton.ParseThroughChildren(container)[0]

		pi := fcSearchSelectors.PageInfo.ParseThroughChildren(container)
		if len(pi) > 1 {
			curPage, err := strconv.Atoi(pi[0])
			if err == nil {
				pageInfo.CurrentPage = curPage
			}
			totalPages, err := strconv.Atoi(pi[1])
			if err == nil {
				pageInfo.TotalPages = totalPages
			}
		}

		container.ForEach(entrySelectors.Root.Selector, func(i int, e *colly.HTMLElement) {
			nextFC := FreeCompanySearchResult{
				Active: FreeCompanyActiveState(entrySelectors.Active.ParseThroughChildren(e)[0]),
				Name:   entrySelectors.Name.ParseThroughChildren(e)[0],
				ID:     entrySelectors.ID.ParseThroughChildren(e)[0],
				CrestLayers: &CrestLayers{
					Bottom: entrySelectors.CrestLayers.Bottom.ParseThroughChildren(e)[0],
					Middle: entrySelectors.CrestLayers.Middle.ParseThroughChildren(e)[0],
					Top:    entrySelectors.CrestLayers.Top.ParseThroughChildren(e)[0],
				},
				Estate:      entrySelectors.EstateBuilt.ParseThroughChildren(e)[0],
				Recruitment: FreeCompanyRecruitingState(entrySelectors.RecruitmentOpen.ParseThroughChildren(e)[0]),
			}

			gcName := entrySelectors.GrandCompany.ParseThroughChildren(e)[0]
			gc := s.grandCompanyTableLookup(gcName)

			nGCs := s.getGrandCompanyTable().GrandCompaniesLength()
			for i := 0; i < nGCs; i++ {
				nextFC.GrandCompany = &NamedEntity{
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

		revisited := false
		if !revisited && nextURI == "" {
			revisited = true
			err := c.Visit(container.Request.URL.String())
			if err != nil {
				output <- &FreeCompanySearchResult{
					Error: err,
				}
			}
		}
	})

	return c
}

func (s *Scraper) buildLinkshellSearchCollector(
	pageInfo *PageInfo,
	output chan *LinkshellSearchResult,
) *colly.Collector {
	c := colly.NewCollector(
		colly.MaxDepth(2),
		colly.UserAgent(s.meta.UserAgentDesktop),
		colly.IgnoreRobotsTxt(),
		colly.AllowURLRevisit(),
		colly.Async(),
	)

	lsSearchSelectors := s.getSearchSelectors().Linkshell
	entrySelectors := lsSearchSelectors.Entry

	c.OnHTML(lsSearchSelectors.Root.Selector, func(container *colly.HTMLElement) {
		nextURI := lsSearchSelectors.ListNextButton.ParseThroughChildren(container)[0]

		pi := lsSearchSelectors.PageInfo.ParseThroughChildren(container)
		if len(pi) > 1 {
			curPage, err := strconv.Atoi(pi[0])
			if err == nil {
				pageInfo.CurrentPage = curPage
			}
			totalPages, err := strconv.Atoi(pi[1])
			if err == nil {
				pageInfo.TotalPages = totalPages
			}
		}

		container.ForEach(entrySelectors.Root.Selector, func(i int, e *colly.HTMLElement) {
			nextLinkshell := LinkshellSearchResult{
				Name: entrySelectors.Name.ParseThroughChildren(e)[0],
				ID:   entrySelectors.ID.ParseThroughChildren(e)[0],
			}

			server := entrySelectors.Server.ParseThroughChildren(e)
			nextLinkshell.World = server[0]
			nextLinkshell.DC = server[1]

			activeMembersStr := entrySelectors.ActiveMembers.ParseThroughChildren(e)[0]
			activeMembers, err := strconv.ParseUint(activeMembersStr, 10, 32)
			if err == nil {
				nextLinkshell.ActiveMembers = uint32(activeMembers)
			}

			output <- &nextLinkshell
		})

		revisited := false
		if !revisited && nextURI == "" {
			revisited = true
			err := c.Visit(container.Request.URL.String())
			if err != nil {
				output <- &LinkshellSearchResult{
					Error: err,
				}
			}
		}
	})

	return c
}

func (s *Scraper) buildPVPTeamSearchCollector(
	pageInfo *PageInfo,
	output chan *PVPTeamSearchResult,
) *colly.Collector {
	c := colly.NewCollector(
		colly.MaxDepth(2),
		colly.UserAgent(s.meta.UserAgentDesktop),
		colly.IgnoreRobotsTxt(),
		colly.AllowURLRevisit(),
		colly.Async(),
	)

	pvpTeamSearchSelectors := s.getSearchSelectors().PVPTeam
	entrySelectors := pvpTeamSearchSelectors.Entry

	c.OnHTML(pvpTeamSearchSelectors.Root.Selector, func(container *colly.HTMLElement) {
		nextURI := pvpTeamSearchSelectors.ListNextButton.ParseThroughChildren(container)[0]

		pi := pvpTeamSearchSelectors.PageInfo.ParseThroughChildren(container)
		if len(pi) > 1 {
			curPage, err := strconv.Atoi(pi[0])
			if err == nil {
				pageInfo.CurrentPage = curPage
			}
			totalPages, err := strconv.Atoi(pi[1])
			if err == nil {
				pageInfo.TotalPages = totalPages
			}
		}

		container.ForEach(entrySelectors.Root.Selector, func(i int, e *colly.HTMLElement) {
			nextTeam := PVPTeamSearchResult{
				Name: entrySelectors.Name.ParseThroughChildren(e)[0],
				ID:   entrySelectors.ID.ParseThroughChildren(e)[0],
				DC:   entrySelectors.DC.ParseThroughChildren(e)[0],
				CrestLayers: &CrestLayers{
					Bottom: entrySelectors.CrestLayers.Bottom.ParseThroughChildren(e)[0],
					Middle: entrySelectors.CrestLayers.Middle.ParseThroughChildren(e)[0],
					Top:    entrySelectors.CrestLayers.Top.ParseThroughChildren(e)[0],
				},
			}

			output <- &nextTeam
		})

		revisited := false
		if !revisited && nextURI == "" {
			revisited = true
			err := c.Visit(container.Request.URL.String())
			if err != nil {
				output <- &PVPTeamSearchResult{
					Error: err,
				}
			}
		}
	})

	return c
}
