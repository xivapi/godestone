package godestone

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/karashiiro/godestone/collectors"
	"github.com/karashiiro/godestone/pack/css"
	"github.com/karashiiro/godestone/pack/exports"
	"github.com/karashiiro/godestone/search"

	"github.com/karashiiro/godestone/models"
	"github.com/karashiiro/godestone/selectors"
)

// Scraper is the object through which interactions with The Lodestone are made.
type Scraper struct {
	lang SiteLang

	meta *models.Meta

	cwlsSelectors      *selectors.CWLSSelectors
	linkshellSelectors *selectors.LinkshellSelectors
	profileSelectors   *selectors.ProfileSelectors
	pvpTeamSelectors   *selectors.PVPTeamSelectors
	searchSelectors    *selectors.SearchSelectors
	fcSelectors        *selectors.FreeCompanySelectors

	achievementTable *exports.AchievementTable
	itemTable        *exports.ItemTable
	minionTable      *exports.MinionTable
	mountTable       *exports.MountTable
	titleTable       *exports.TitleTable
}

// NewScraper creates a new instance of the Scraper.
func NewScraper(lang SiteLang) *Scraper {
	metaBytes, _ := css.Asset("meta.json")
	meta := models.Meta{}
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

func (s *Scraper) getTitleTable() *exports.TitleTable {
	if s.titleTable == nil {
		data, _ := exports.Asset("title_table.bin")
		titleTable := exports.GetRootAsTitleTable(data, 0)
		s.titleTable = titleTable
	}
	return s.titleTable
}

// FetchCharacter returns character information for the provided Lodestone ID.
func (s *Scraper) FetchCharacter(id uint32) (*models.Character, error) {
	now := time.Now()
	charData := models.Character{ID: id, ParseDate: now}

	charCollector := collectors.BuildCharacterCollector(s.meta, s.getProfileSelectors(), s.getTitleTable(), &charData)
	err := charCollector.Visit(fmt.Sprintf("https://%s.finalfantasyxiv.com/lodestone/character/%d", s.lang, id))
	if err != nil {
		return nil, err
	}
	charCollector.Wait()

	classJobCollector := collectors.BuildClassJobCollector(s.meta, s.getProfileSelectors(), &charData)
	err = classJobCollector.Visit(fmt.Sprintf("https://%s.finalfantasyxiv.com/lodestone/character/%d/class_job/", s.lang, id))
	if err != nil {
		return nil, err
	}
	classJobCollector.Wait()

	return &charData, nil
}

// FetchCharacterMinions returns unlocked minion information for the provided Lodestone ID.
func (s *Scraper) FetchCharacterMinions(id uint32) ([]*models.Minion, error) {
	output := make(chan *models.Minion)
	errors := make(chan error, 1)
	done := make(chan bool, 1)

	go func() {
		minionCollector := collectors.BuildMinionCollector(s.meta, s.getProfileSelectors(), s.getMinionTable(), output)

		err := minionCollector.Visit(fmt.Sprintf("https://%s.finalfantasyxiv.com/lodestone/character/%d/minion/", s.lang, id))
		if err != nil && err.Error() != http.StatusText(http.StatusNotFound) {
			errors <- err
		}

		minionCollector.Wait()

		close(output)
		close(errors)

		done <- true
		close(done)
	}()

	minions := []*models.Minion{}
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

// FetchCharacterMounts returns unlocked mount information for the provided Lodestone ID.
func (s *Scraper) FetchCharacterMounts(id uint32) ([]*models.Mount, error) {
	output := make(chan *models.Mount)
	errors := make(chan error, 1)
	done := make(chan bool, 1)

	go func() {
		mountCollector := collectors.BuildMountCollector(s.meta, s.getProfileSelectors(), s.getMountTable(), output)

		err := mountCollector.Visit(fmt.Sprintf("https://%s.finalfantasyxiv.com/lodestone/character/%d/mount/", s.lang, id))
		if err != nil && err.Error() != http.StatusText(http.StatusNotFound) {
			errors <- err
		}

		mountCollector.Wait()

		close(output)
		close(errors)

		done <- true
		close(done)
	}()

	mounts := []*models.Mount{}
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

// FetchCharacterAchievements returns unlocked achievement information for the provided Lodestone ID.
func (s *Scraper) FetchCharacterAchievements(id uint32) chan *models.AchievementInfo {
	output := make(chan *models.AchievementInfo)

	go func() {
		achievementCollector := collectors.BuildAchievementCollector(s.meta, s.getProfileSelectors(), s.getAchievementTable(), output)
		err := achievementCollector.Visit(fmt.Sprintf("https://%s.finalfantasyxiv.com/lodestone/character/%d/achievement/", s.lang, id))
		if err != nil {
			aai := &models.AllAchievementInfo{}
			errAi := &models.AchievementInfo{
				AllAchievementInfo: aai,
				Error:              err,
			}

			if err.Error() == http.StatusText(http.StatusForbidden) {
				aai.Private = true
			}

			output <- errAi
		}
		achievementCollector.Wait()

		close(output)
	}()

	return output
}

// FetchLinkshell returns linkshell information for the provided linkshell ID.
func (s *Scraper) FetchLinkshell(id string) (*models.Linkshell, error) {
	now := time.Now()
	ls := models.Linkshell{ID: id, ParseDate: now}

	lsCollector := collectors.BuildLinkshellCollector(s.meta, s.getLinkshellSelectors(), &ls)
	err := lsCollector.Visit(fmt.Sprintf("https://%s.finalfantasyxiv.com/lodestone/linkshell/%s", s.lang, id))
	if err != nil {
		return nil, err
	}
	lsCollector.Wait()

	return &ls, nil
}

// FetchCWLS returns CWLS information for the provided CWLS ID.
func (s *Scraper) FetchCWLS(id string) (*models.CWLS, error) {
	now := time.Now()
	cwls := models.CWLS{ID: id, ParseDate: now}

	cwlsCollector := collectors.BuildCWLSCollector(s.meta, s.getCWLSSelectors(), &cwls)
	err := cwlsCollector.Visit(fmt.Sprintf("https://%s.finalfantasyxiv.com/lodestone/crossworld_linkshell/%s", s.lang, id))
	if err != nil {
		return nil, err
	}
	cwlsCollector.Wait()

	return &cwls, nil
}

// FetchPVPTeam returns PVP team information for the provided PVP team ID.
func (s *Scraper) FetchPVPTeam(id string) (*models.PVPTeam, error) {
	now := time.Now()
	pvpTeam := models.PVPTeam{ID: id, ParseDate: now}

	pvpTeamCollector := collectors.BuildPVPTeamCollector(s.meta, s.getPVPTeamSelectors(), &pvpTeam)
	err := pvpTeamCollector.Visit(fmt.Sprintf("https://%s.finalfantasyxiv.com/lodestone/pvpteam/%s", s.lang, id))
	if err != nil {
		return nil, err
	}
	pvpTeamCollector.Wait()

	return &pvpTeam, nil
}

// FetchFreeCompany returns Free Company information for the provided Free Company ID.
func (s *Scraper) FetchFreeCompany(id string) (*models.FreeCompany, error) {
	now := time.Now()
	fc := models.FreeCompany{ID: id, ParseDate: now}

	fcCollector := collectors.BuildFreeCompanyCollector(s.meta, s.getFreeCompanySelectors(), &fc)
	err := fcCollector.Visit(fmt.Sprintf("https://%s.finalfantasyxiv.com/lodestone/freecompany/%s", s.lang, id))
	if err != nil {
		return nil, err
	}
	fcCollector.Wait()

	return &fc, nil
}

// FetchFreeCompanyMembers returns Free Company member information for the provided Free Company ID.
func (s *Scraper) FetchFreeCompanyMembers(id string) chan *models.FreeCompanyMember {
	output := make(chan *models.FreeCompanyMember)

	go func() {
		fcMembersCollector := collectors.BuildFreeCompanyMembersCollector(s.meta, s.getFreeCompanySelectors(), output)

		err := fcMembersCollector.Visit(fmt.Sprintf("https://%s.finalfantasyxiv.com/lodestone/freecompany/%s/member/", s.lang, id))
		if err != nil {
			output <- &models.FreeCompanyMember{
				Error: err,
			}
		}
		fcMembersCollector.Wait()

		close(output)
	}()

	return output
}

// SearchFreeCompanies returns a channel of searchable Free Companies.
func (s *Scraper) SearchFreeCompanies(opts search.FreeCompanyOptions) chan *models.FreeCompanySearchResult {
	output := make(chan *models.FreeCompanySearchResult)

	uri := opts.BuildURI(string(s.lang))
	go func() {
		searchCollector := collectors.BuildFreeCompanySearchCollector(s.meta, s.getSearchSelectors(), output)

		err := searchCollector.Visit(uri)
		if err != nil {
			output <- &models.FreeCompanySearchResult{
				Error: err,
			}
		}
		searchCollector.Wait()

		close(output)
	}()

	return output
}

// SearchCharacters returns a channel of searchable characters.
func (s *Scraper) SearchCharacters(opts search.CharacterOptions) chan *models.CharacterSearchResult {
	output := make(chan *models.CharacterSearchResult)

	uri := opts.BuildURI(string(s.lang))
	go func() {
		searchCollector := collectors.BuildCharacterSearchCollector(s.meta, s.getSearchSelectors(), output)

		err := searchCollector.Visit(uri)
		if err != nil {
			output <- &models.CharacterSearchResult{
				Error: err,
			}
		}
		searchCollector.Wait()

		close(output)
	}()

	return output
}

// SearchCWLS returns a channel of searchable crossworld linkshells.
func (s *Scraper) SearchCWLS(opts search.CWLSOptions) chan *models.CWLSSearchResult {
	output := make(chan *models.CWLSSearchResult)

	uri := opts.BuildURI(string(s.lang))
	go func() {
		searchCollector := collectors.BuildCWLSSearchCollector(s.meta, s.getSearchSelectors(), output)

		err := searchCollector.Visit(uri)
		if err != nil {
			output <- &models.CWLSSearchResult{
				Error: err,
			}
		}
		searchCollector.Wait()

		close(output)
	}()

	return output
}

// SearchLinkshells returns a channel of searchable linkshells.
func (s *Scraper) SearchLinkshells(opts search.LinkshellOptions) chan *models.LinkshellSearchResult {
	output := make(chan *models.LinkshellSearchResult)

	uri := opts.BuildURI(string(s.lang))
	go func() {
		searchCollector := collectors.BuildLinkshellSearchCollector(s.meta, s.getSearchSelectors(), output)

		err := searchCollector.Visit(uri)
		if err != nil {
			output <- &models.LinkshellSearchResult{
				Error: err,
			}
		}
		searchCollector.Wait()

		close(output)
	}()

	return output
}

// SearchPVPTeams returns a channel of searchable PVP teams.
func (s *Scraper) SearchPVPTeams(opts search.PVPTeamOptions) chan *models.PVPTeamSearchResult {
	output := make(chan *models.PVPTeamSearchResult)

	uri := opts.BuildURI(string(s.lang))
	go func() {
		searchCollector := collectors.BuildPVPTeamSearchCollector(s.meta, s.getSearchSelectors(), output)

		err := searchCollector.Visit(uri)
		if err != nil {
			output <- &models.PVPTeamSearchResult{
				Error: err,
			}
		}
		searchCollector.Wait()

		close(output)
	}()

	return output
}
