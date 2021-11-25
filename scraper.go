package godestone

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/xivapi/godestone/v2/internal/pack/css"
	"github.com/xivapi/godestone/v2/internal/selectors"
	"github.com/xivapi/godestone/v2/provider"
)

// Used for a band-aid in the Character scraper
var elementalLevelNames = map[string]struct{}{
	"Elemental Level":    {},
	"Elementarstufe":     {},
	"Niveau élémentaire": {},
	"エレメンタルレベル":          {},
}

// Scraper is the object through which interactions with The Lodestone are made.
type Scraper struct {
	lang SiteLang

	meta *meta

	dataProvider provider.DataProvider

	cwlsSelectors      *selectors.CWLSSelectors
	linkshellSelectors *selectors.LinkshellSelectors
	profileSelectors   *selectors.ProfileSelectors
	pvpTeamSelectors   *selectors.PVPTeamSelectors
	searchSelectors    *selectors.SearchSelectors
	fcSelectors        *selectors.FreeCompanySelectors
}

// NewScraper creates a new instance of the Scraper. Do note that all five language-versions of the website
// are on the same physical servers in Japan. Changing the language of the website will not meaningfully
// improve response times.
func NewScraper(dataProvider provider.DataProvider, lang SiteLang) *Scraper {
	metaBytes, _ := css.Asset("meta.json")
	meta := meta{}
	json.Unmarshal(metaBytes, &meta)

	return &Scraper{
		lang:         lang,
		meta:         &meta,
		dataProvider: dataProvider,
	}
}

func (s *Scraper) getCWLSSelectors() *selectors.CWLSSelectors {
	if s.cwlsSelectors == nil {
		s.cwlsSelectors = selectors.LoadCWLSSelectors()
	}
	return s.cwlsSelectors
}

func (s *Scraper) getLinkshellSelectors() *selectors.LinkshellSelectors {
	if s.linkshellSelectors == nil {
		s.linkshellSelectors = selectors.LoadLinkshellSelectors()
	}
	return s.linkshellSelectors
}

func (s *Scraper) getProfileSelectors() *selectors.ProfileSelectors {
	if s.profileSelectors == nil {
		s.profileSelectors = selectors.LoadProfileSelectors()
	}
	return s.profileSelectors
}

func (s *Scraper) getPVPTeamSelectors() *selectors.PVPTeamSelectors {
	if s.pvpTeamSelectors == nil {
		s.pvpTeamSelectors = selectors.LoadPVPTeamSelectors()
	}
	return s.pvpTeamSelectors
}

func (s *Scraper) getSearchSelectors() *selectors.SearchSelectors {
	if s.searchSelectors == nil {
		s.searchSelectors = selectors.LoadSearchSelectors()
	}
	return s.searchSelectors
}

func (s *Scraper) getFreeCompanySelectors() *selectors.FreeCompanySelectors {
	if s.fcSelectors == nil {
		s.fcSelectors = selectors.LoadFreeCompanySelectors()
	}
	return s.fcSelectors
}

// FetchCharacter returns character information for the provided Lodestone ID. This function makes
// two requests: one to the base character profile, and another to the class and job page, returning
// an error if either request fails.
func (s *Scraper) FetchCharacter(id uint32) (*Character, error) {
	now := time.Now()
	charData := Character{ID: id, ParseDate: now}
	errors := make(chan error, 2)

	charCollector := s.buildCharacterCollector(&charData)

	charCollector.OnError(func(r *colly.Response, err error) {
		errors <- err
	})

	charSelectors := s.getProfileSelectors().Character

	icon := ""
	charCollector.OnHTML(charSelectors.ActiveClassJob.Selector, func(e *colly.HTMLElement) {
		icon = charSelectors.ActiveClassJob.Parse(e)[0]
	})

	activeClassJobName := ""
	charCollector.OnHTML(charSelectors.ClassJobIcons.Root.Selector, func(container *colly.HTMLElement) {
		container.ForEach(charSelectors.ClassJobIcons.Icon.Selector, func(i int, e *colly.HTMLElement) {
			thisIcon := charSelectors.ClassJobIcons.Icon.ParseThroughChildren(container)[0]
			if icon == thisIcon {
				activeClassJobName = e.Attr("data-tooltip")
			}
		})
	})
	charCollector.Visit(fmt.Sprintf("https://%s.finalfantasyxiv.com/lodestone/character/%d", s.lang, id))

	classJobCollector := s.buildClassJobCollector(&charData)
	classJobCollector.Visit(fmt.Sprintf("https://%s.finalfantasyxiv.com/lodestone/character/%d/class_job/", s.lang, id))

	charCollector.Wait()
	classJobCollector.Wait()
	close(errors)
	select {
	case err, ok := <-errors:
		if ok {
			return nil, err
		}
	}

	// Link the active classjob details from the classjob page
	for _, cj := range charData.ClassJobs {
		if cj.Name == activeClassJobName {
			charData.ActiveClassJob = cj
		}
	}

	// TODO: https://github.com/xivapi/godestone/issues/17 Ugly band-aid
	// If the Bozja struct has Eureka data, we'll just move the data
	// over and clear the Bozja struct for the time being. This is
	// implemented as unintrusively as possible so that we can rip it
	// out as soon as a better solution becomes available.
	if _, ok := elementalLevelNames[charData.ClassJobBozjan.Name]; ok {
		charData.ClassJobElemental.Level = charData.ClassJobBozjan.Level
		charData.ClassJobElemental.Name = charData.ClassJobBozjan.Name

		expStrs := s.profileSelectors.ClassJob.Eureka.Exp.Parse(charData.ClassJobBozjan.mettleRaw)
		expStrs[0] = nonDigits.ReplaceAllString(expStrs[0], "")
		expStrs[1] = nonDigits.ReplaceAllString(expStrs[1], "")

		curExp, err := strconv.ParseUint(expStrs[0], 10, 32)
		if err == nil {
			charData.ClassJobElemental.ExpLevel = uint32(curExp)
		}

		maxExp, err := strconv.ParseUint(expStrs[1], 10, 32)
		if err == nil {
			charData.ClassJobElemental.ExpLevelMax = uint32(maxExp)
		}

		charData.ClassJobElemental.ExpLevelTogo = charData.ClassJobElemental.ExpLevelMax - charData.ClassJobElemental.ExpLevel

		// Clear old data
		charData.ClassJobBozjan.Level = 0
		charData.ClassJobBozjan.Mettle = 0
		charData.ClassJobBozjan.Name = ""
	}

	return &charData, nil
}

// FetchCharacterMinions returns unlocked minion information for the provided Lodestone ID. The error is returned
// if the request fails with anything other than a 404. A 404 can result from a character not existing, but it can
// also result from a character not having any minions.
func (s *Scraper) FetchCharacterMinions(id uint32) ([]*Minion, error) {
	output := make(chan *Minion)
	errors := make(chan error, 1)
	done := make(chan bool, 1)

	go func() {
		minionCollector := s.buildMinionCollector(output)
		minionCollector.OnError(func(r *colly.Response, err error) {
			if err.Error() != http.StatusText(http.StatusNotFound) {
				errors <- err
			}
		})
		minionCollector.Visit(fmt.Sprintf("https://%s.finalfantasyxiv.com/lodestone/character/%d/minion/", s.lang, id))
		minionCollector.Wait()

		close(output)
		close(errors)

		done <- true
		close(done)
	}()

	minions := make([]*Minion, 0)
	for minion := range output {
		minions = append(minions, minion)
	}

	<-done
	select {
	case err, ok := <-errors:
		if ok {
			return nil, err
		}
	}
	return minions, nil
}

// FetchCharacterMounts returns unlocked mount information for the provided Lodestone ID. The error is returned
// if the request fails with anything other than a 404. A 404 can result from a character not existing, but it can
// also result from a character not having any mounts.
func (s *Scraper) FetchCharacterMounts(id uint32) ([]*Mount, error) {
	output := make(chan *Mount)
	errors := make(chan error, 1)
	done := make(chan bool, 1)

	go func() {
		mountCollector := s.buildMountCollector(output)
		mountCollector.OnError(func(r *colly.Response, err error) {
			if err.Error() != http.StatusText(http.StatusNotFound) {
				errors <- err
			}
		})
		mountCollector.Visit(fmt.Sprintf("https://%s.finalfantasyxiv.com/lodestone/character/%d/mount/", s.lang, id))
		mountCollector.Wait()

		close(output)
		close(errors)

		done <- true
		close(done)
	}()

	mounts := make([]*Mount, 0)
	for mount := range output {
		mounts = append(mounts, mount)
	}

	<-done
	select {
	case err, ok := <-errors:
		if ok {
			return nil, err
		}
	}
	return mounts, nil
}

// FetchCharacterAchievements returns unlocked achievement information for the provided Lodestone ID. The error
// is returned if the request fails with anything other than a 403. A 403 will not be raised when the character's
// achievements are private. Instead, the Private field on the AllAchievementInfo object will be set to true.
func (s *Scraper) FetchCharacterAchievements(id uint32) ([]*AchievementInfo, *AllAchievementInfo, error) {
	output := make(chan *AchievementInfo)
	errors := make(chan error, 1)
	done := make(chan bool, 1)

	allAchievementInfo := &AllAchievementInfo{}

	go func() {
		achievementCollector := s.buildAchievementCollector(allAchievementInfo, output, errors)
		achievementCollector.OnError(func(r *colly.Response, err error) {
			if err.Error() != http.StatusText(http.StatusForbidden) {
				errors <- err
			} else {
				// 403
				allAchievementInfo.Private = true
			}
		})
		achievementCollector.Visit(fmt.Sprintf("https://%s.finalfantasyxiv.com/lodestone/character/%d/achievement/", s.lang, id))
		achievementCollector.Wait()

		close(output)
		close(errors)

		done <- true
		close(done)
	}()

	achievements := make([]*AchievementInfo, 0)
	for achievement := range output {
		achievements = append(achievements, achievement)
	}

	<-done
	select {
	case err, ok := <-errors:
		if ok {
			return nil, nil, err
		}
	}
	return achievements, allAchievementInfo, nil
}

// FetchLinkshell returns linkshell information for the provided linkshell ID. The error is returned if the
// request fails.
func (s *Scraper) FetchLinkshell(id string) (*Linkshell, error) {
	now := time.Now()
	ls := Linkshell{ID: id, ParseDate: now}
	errors := make(chan error, 1)

	lsCollector := s.buildLinkshellCollector(&ls)
	lsCollector.OnError(func(r *colly.Response, err error) {
		errors <- err
	})
	lsCollector.Visit(fmt.Sprintf("https://%s.finalfantasyxiv.com/lodestone/linkshell/%s", s.lang, id))
	lsCollector.Wait()
	close(errors)
	select {
	case err, ok := <-errors:
		if ok {
			return nil, err
		}
	}

	return &ls, nil
}

// FetchCWLS returns CWLS information for the provided CWLS ID. The error is returned if the
// request fails.
func (s *Scraper) FetchCWLS(id string) (*CWLS, error) {
	now := time.Now()
	cwls := CWLS{ID: id, ParseDate: now}
	errors := make(chan error, 1)

	cwlsCollector := s.buildCWLSCollector(&cwls)
	cwlsCollector.OnError(func(r *colly.Response, err error) {
		errors <- err
	})
	cwlsCollector.Visit(fmt.Sprintf("https://%s.finalfantasyxiv.com/lodestone/crossworld_linkshell/%s", s.lang, id))
	cwlsCollector.Wait()
	close(errors)
	select {
	case err, ok := <-errors:
		if ok {
			return nil, err
		}
	}

	return &cwls, nil
}

// FetchPVPTeam returns PVP team information for the provided PVP team ID. The error is returned if the
// request fails.
func (s *Scraper) FetchPVPTeam(id string) (*PVPTeam, error) {
	now := time.Now()
	pvpTeam := PVPTeam{ID: id, ParseDate: now}
	errors := make(chan error, 1)

	pvpTeamCollector := s.buildPVPTeamCollector(&pvpTeam)
	pvpTeamCollector.OnError(func(r *colly.Response, err error) {
		errors <- err
	})
	pvpTeamCollector.Visit(fmt.Sprintf("https://%s.finalfantasyxiv.com/lodestone/pvpteam/%s", s.lang, id))
	pvpTeamCollector.Wait()
	close(errors)
	select {
	case err, ok := <-errors:
		if ok {
			return nil, err
		}
	}

	return &pvpTeam, nil
}

// FetchFreeCompany returns Free Company information for the provided Free Company ID. The error is returned if the
// request fails.
func (s *Scraper) FetchFreeCompany(id string) (*FreeCompany, error) {
	now := time.Now()
	fc := FreeCompany{ID: id, ParseDate: now}
	errors := make(chan error, 1)

	fcCollector := s.buildFreeCompanyCollector(&fc)
	fcCollector.OnError(func(r *colly.Response, err error) {
		errors <- err
	})
	fcCollector.Visit(fmt.Sprintf("https://%s.finalfantasyxiv.com/lodestone/freecompany/%s", s.lang, id))
	fcCollector.Wait()
	close(errors)
	select {
	case err, ok := <-errors:
		if ok {
			return nil, err
		}
	}

	return &fc, nil
}

// FetchFreeCompanyMembers returns Free Company member information for the provided Free Company ID. The error
// is returned if the request fails.
func (s *Scraper) FetchFreeCompanyMembers(id string) chan *FreeCompanyMember {
	output := make(chan *FreeCompanyMember)

	go func() {
		fcMembersCollector := s.buildFreeCompanyMembersCollector(output)
		fcMembersCollector.OnError(func(r *colly.Response, err error) {
			output <- &FreeCompanyMember{
				Error: err,
			}
		})
		fcMembersCollector.Visit(fmt.Sprintf("https://%s.finalfantasyxiv.com/lodestone/freecompany/%s/member/", s.lang, id))
		fcMembersCollector.Wait()
		close(output)
	}()

	return output
}

// SearchFreeCompanies returns a channel of searchable Free Companies. Please note that searches are notoriously
// poor, and often return exact matches far down in the results, or else return no search results when search
// results should be present. This library does one retry on each failure, but this is not a guarantee that
// all search results will be returned.
func (s *Scraper) SearchFreeCompanies(opts FreeCompanyOptions) chan *FreeCompanySearchResult {
	output := make(chan *FreeCompanySearchResult)

	uri := opts.BuildURI(string(s.lang))
	go func() {
		pageInfo := &PageInfo{TotalPages: 20}
		revisited := map[string]bool{}

		mu := sync.Mutex{}

		searchCollector := s.buildFreeCompanySearchCollector(
			pageInfo,
			output,
		)
		searchCollector.OnError(func(r *colly.Response, err error) {
			url := r.Request.URL.String()
			mu.Lock()
			if revisited[url] {
				output <- &FreeCompanySearchResult{
					Error: err,
				}
			} else {
				searchCollector.Visit(url)
			}
			revisited[url] = true
			mu.Unlock()
		})
		searchCollector.Visit(uri)
		searchCollector.Wait()

		done := make(chan bool, pageInfo.TotalPages-1)
		for i := 2; i <= pageInfo.TotalPages; i++ {
			nextURI := uri + fmt.Sprintf("&page=%d", i)
			go func() {
				searchCollector.Visit(nextURI)
				searchCollector.Wait()
				done <- true
			}()
		}

		for i := 2; i <= pageInfo.TotalPages; i++ {
			<-done
		}

		close(output)
	}()

	return output
}

// SearchCharacters returns a channel of searchable characters. Please note that searches are notoriously
// poor, and often return exact matches far down in the results, or else return no search results when search
// results should be present. This library does one retry on each failure, but this is not a guarantee that
// all search results will be returned.
func (s *Scraper) SearchCharacters(opts CharacterOptions) chan *CharacterSearchResult {
	output := make(chan *CharacterSearchResult)

	uri := opts.BuildURI(s, string(s.lang))

	go func() {
		pageInfo := &PageInfo{TotalPages: 20}
		revisited := map[string]bool{}

		mu := sync.Mutex{}

		searchCollector := s.buildCharacterSearchCollector(
			pageInfo,
			output,
		)
		searchCollector.OnError(func(r *colly.Response, err error) {
			url := r.Request.URL.String()
			mu.Lock()
			if revisited[url] {
				output <- &CharacterSearchResult{
					Error: err,
				}
			} else {
				searchCollector.Visit(url)
			}
			revisited[url] = true
			mu.Unlock()
		})
		searchCollector.Visit(uri)
		searchCollector.Wait()

		done := make(chan bool, pageInfo.TotalPages-1)
		for i := 2; i <= pageInfo.TotalPages; i++ {
			nextURI := uri + fmt.Sprintf("&page=%d", i)
			go func() {
				searchCollector.Visit(nextURI)
				searchCollector.Wait()
				done <- true
			}()
		}

		for i := 2; i <= pageInfo.TotalPages; i++ {
			<-done
		}

		close(output)
	}()

	return output
}

// SearchCWLS returns a channel of searchable crossworld linkshells. Please note that searches are notoriously
// poor, and often return exact matches far down in the results, or else return no search results when search
// results should be present. This library does one retry on each failure, but this is not a guarantee that
// all search results will be returned.
func (s *Scraper) SearchCWLS(opts CWLSOptions) chan *CWLSSearchResult {
	output := make(chan *CWLSSearchResult)

	uri := opts.BuildURI(string(s.lang))
	go func() {
		pageInfo := &PageInfo{TotalPages: 20}
		revisited := map[string]bool{}

		mu := sync.Mutex{}

		searchCollector := s.buildCWLSSearchCollector(
			pageInfo,
			output,
		)
		searchCollector.OnError(func(r *colly.Response, err error) {
			url := r.Request.URL.String()
			mu.Lock()
			if revisited[url] {
				output <- &CWLSSearchResult{
					Error: err,
				}
			} else {
				searchCollector.Visit(url)
			}
			revisited[url] = true
			mu.Unlock()
		})
		searchCollector.Visit(uri)
		searchCollector.Wait()

		done := make(chan bool, pageInfo.TotalPages-1)
		for i := 2; i <= pageInfo.TotalPages; i++ {
			nextURI := uri + fmt.Sprintf("&page=%d", i)
			go func() {
				searchCollector.Visit(nextURI)
				searchCollector.Wait()
				done <- true
			}()
		}

		for i := 2; i <= pageInfo.TotalPages; i++ {
			<-done
		}

		close(output)
	}()

	return output
}

// SearchLinkshells returns a channel of searchable linkshells. Please note that searches are notoriously
// poor, and often return exact matches far down in the results, or else return no search results when search
// results should be present. This library does one retry on each failure, but this is not a guarantee that
// all search results will be returned.
func (s *Scraper) SearchLinkshells(opts LinkshellOptions) chan *LinkshellSearchResult {
	output := make(chan *LinkshellSearchResult)

	uri := opts.BuildURI(string(s.lang))
	go func() {
		pageInfo := &PageInfo{TotalPages: 20}
		revisited := map[string]bool{}

		mu := sync.Mutex{}

		searchCollector := s.buildLinkshellSearchCollector(
			pageInfo,
			output,
		)
		searchCollector.OnError(func(r *colly.Response, err error) {
			url := r.Request.URL.String()
			mu.Lock()
			if revisited[url] {
				output <- &LinkshellSearchResult{
					Error: err,
				}
			} else {
				searchCollector.Visit(url)
			}
			revisited[url] = true
			mu.Unlock()
		})
		searchCollector.Visit(uri)
		searchCollector.Wait()

		done := make(chan bool, pageInfo.TotalPages-1)
		for i := 2; i <= pageInfo.TotalPages; i++ {
			nextURI := uri + fmt.Sprintf("&page=%d", i)
			go func() {
				searchCollector.Visit(nextURI)
				searchCollector.Wait()
				done <- true
			}()
		}

		for i := 2; i <= pageInfo.TotalPages; i++ {
			<-done
		}

		close(output)
	}()

	return output
}

// SearchPVPTeams returns a channel of searchable PVP teams. Please note that searches are notoriously
// poor, and often return exact matches far down in the results, or else return no search results when search
// results should be present. This library does one retry on each failure, but this is not a guarantee that
// all search results will be returned.
func (s *Scraper) SearchPVPTeams(opts PVPTeamOptions) chan *PVPTeamSearchResult {
	output := make(chan *PVPTeamSearchResult)

	uri := opts.BuildURI(string(s.lang))
	go func() {
		pageInfo := &PageInfo{TotalPages: 20}
		revisited := map[string]bool{}

		mu := sync.Mutex{}

		searchCollector := s.buildPVPTeamSearchCollector(
			pageInfo,
			output,
		)
		searchCollector.OnError(func(r *colly.Response, err error) {
			url := r.Request.URL.String()
			mu.Lock()
			if revisited[url] {
				output <- &PVPTeamSearchResult{
					Error: err,
				}
			} else {
				searchCollector.Visit(url)
			}
			revisited[url] = true
			mu.Unlock()
		})
		searchCollector.Visit(uri)
		searchCollector.Wait()

		done := make(chan bool, pageInfo.TotalPages-1)
		for i := 2; i <= pageInfo.TotalPages; i++ {
			nextURI := uri + fmt.Sprintf("&page=%d", i)
			go func() {
				searchCollector.Visit(nextURI)
				searchCollector.Wait()
				done <- true
			}()
		}

		for i := 2; i <= pageInfo.TotalPages; i++ {
			<-done
		}

		close(output)
	}()

	return output
}
