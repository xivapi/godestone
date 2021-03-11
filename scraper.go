package godestone

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/karashiiro/godestone/pack/css"
	"github.com/karashiiro/godestone/pack/exports"

	"github.com/karashiiro/godestone/selectors"
)

// Scraper is the object through which interactions with The Lodestone are made.
type Scraper struct {
	lang SiteLang

	meta *meta

	cwlsSelectors      *selectors.CWLSSelectors
	linkshellSelectors *selectors.LinkshellSelectors
	profileSelectors   *selectors.ProfileSelectors
	pvpTeamSelectors   *selectors.PVPTeamSelectors
	searchSelectors    *selectors.SearchSelectors
	fcSelectors        *selectors.FreeCompanySelectors

	achievementTable  *exports.AchievementTable
	classJobTable     *exports.ClassJobTable
	deityTable        *exports.DeityTable
	grandCompanyTable *exports.GrandCompanyTable
	itemTable         *exports.ItemTable
	minionTable       *exports.MinionTable
	mountTable        *exports.MountTable
	raceTable         *exports.RaceTable
	repTable          *exports.ReputationTable
	titleTable        *exports.TitleTable
	townTable         *exports.TownTable
	tribeTable        *exports.TribeTable
}

// NewScraper creates a new instance of the Scraper. Do note that all five language-versions of the website
// are on the same physical servers in Japan. Changing the language of the website will not meaningfully
// improve response times.
func NewScraper(lang SiteLang) *Scraper {
	metaBytes, _ := css.Asset("meta.json")
	meta := meta{}
	json.Unmarshal(metaBytes, &meta)

	return &Scraper{
		lang: lang,
		meta: &meta,
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

func (s *Scraper) getAchievementTable() *exports.AchievementTable {
	if s.achievementTable == nil {
		data, _ := exports.Asset("achievement_table.bin")
		achievementTable := exports.GetRootAsAchievementTable(data, 0)
		s.achievementTable = achievementTable
	}
	return s.achievementTable
}

func (s *Scraper) getClassJobTable() *exports.ClassJobTable {
	if s.classJobTable == nil {
		data, _ := exports.Asset("classjob_table.bin")
		classJobTable := exports.GetRootAsClassJobTable(data, 0)
		s.classJobTable = classJobTable
	}
	return s.classJobTable
}

func (s *Scraper) getDeityTable() *exports.DeityTable {
	if s.deityTable == nil {
		data, _ := exports.Asset("deity_table.bin")
		deityTable := exports.GetRootAsDeityTable(data, 0)
		s.deityTable = deityTable
	}
	return s.deityTable
}

func (s *Scraper) getGrandCompanyTable() *exports.GrandCompanyTable {
	if s.grandCompanyTable == nil {
		data, _ := exports.Asset("gc_table.bin")
		grandCompanyTable := exports.GetRootAsGrandCompanyTable(data, 0)
		s.grandCompanyTable = grandCompanyTable
	}
	return s.grandCompanyTable
}

func (s *Scraper) getItemTable() *exports.ItemTable {
	if s.itemTable == nil {
		data, _ := exports.Asset("item_table.bin")
		itemTable := exports.GetRootAsItemTable(data, 0)
		s.itemTable = itemTable
	}
	return s.itemTable
}

func (s *Scraper) getMinionTable() *exports.MinionTable {
	if s.minionTable == nil {
		data, _ := exports.Asset("minion_table.bin")
		minionTable := exports.GetRootAsMinionTable(data, 0)
		s.minionTable = minionTable
	}
	return s.minionTable
}

func (s *Scraper) getMountTable() *exports.MountTable {
	if s.mountTable == nil {
		data, _ := exports.Asset("mount_table.bin")
		mountTable := exports.GetRootAsMountTable(data, 0)
		s.mountTable = mountTable
	}
	return s.mountTable
}

func (s *Scraper) getRaceTable() *exports.RaceTable {
	if s.raceTable == nil {
		data, _ := exports.Asset("race_table.bin")
		raceTable := exports.GetRootAsRaceTable(data, 0)
		s.raceTable = raceTable
	}
	return s.raceTable
}

func (s *Scraper) getReputationTable() *exports.ReputationTable {
	if s.repTable == nil {
		data, _ := exports.Asset("reputation_table.bin")
		repTable := exports.GetRootAsReputationTable(data, 0)
		s.repTable = repTable
	}
	return s.repTable
}

func (s *Scraper) getTitleTable() *exports.TitleTable {
	if s.titleTable == nil {
		data, _ := exports.Asset("title_table.bin")
		titleTable := exports.GetRootAsTitleTable(data, 0)
		s.titleTable = titleTable
	}
	return s.titleTable
}

func (s *Scraper) getTownTable() *exports.TownTable {
	if s.townTable == nil {
		data, _ := exports.Asset("town_table.bin")
		townTable := exports.GetRootAsTownTable(data, 0)
		s.townTable = townTable
	}
	return s.townTable
}

func (s *Scraper) getTribeTable() *exports.TribeTable {
	if s.tribeTable == nil {
		data, _ := exports.Asset("tribe_table.bin")
		tribeTable := exports.GetRootAsTribeTable(data, 0)
		s.tribeTable = tribeTable
	}
	return s.tribeTable
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

	for _, cj := range charData.ClassJobs {
		if cj.Name == activeClassJobName {
			charData.ActiveClassJob = cj
		}
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
// is returned if the request fails with anything other than a 403. A 403 will be raised when the character's
// achievements are private.
func (s *Scraper) FetchCharacterAchievements(id uint32) ([]*AchievementInfo, *AllAchievementInfo, error) {
	output := make(chan *AchievementInfo)
	errors := make(chan error, 1)
	done := make(chan bool, 1)

	allAchievementInfo := &AllAchievementInfo{}

	go func() {
		achievementCollector := s.buildAchievementCollector(allAchievementInfo, output, errors)
		achievementCollector.OnError(func(r *colly.Response, err error) {
			if err.Error() != http.StatusText(http.StatusNotFound) {
				errors <- err
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
