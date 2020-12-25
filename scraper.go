package godestone

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/karashiiro/godestone/collectors"
	"github.com/karashiiro/godestone/search"

	"github.com/karashiiro/godestone/models"
	"github.com/karashiiro/godestone/pack"
	"github.com/karashiiro/godestone/selectors"
)

// Scraper is the object through which interactions with The Lodestone are made.
type Scraper struct {
	lang SiteLang

	meta               *models.Meta
	cwlsSelectors      *selectors.CWLSSelectors
	fcSelectors        *selectors.FreeCompanySelectors
	linkshellSelectors *selectors.LinkshellSelectors
	profileSelectors   *selectors.ProfileSelectors
	pvpTeamSelectors   *selectors.PVPTeamSelectors
	searchSelectors    *selectors.SearchSelectors
}

// FetchCharacter returns character information for the provided Lodestone ID.
func (s *Scraper) FetchCharacter(id uint32) (*models.Character, error) {
	now := time.Now()
	charData := models.Character{ID: id, ParseDate: now}

	charCollector := collectors.BuildCharacterCollector(s.meta, s.profileSelectors, &charData)
	err := charCollector.Visit(fmt.Sprintf("https://%s.finalfantasyxiv.com/lodestone/character/%d", s.lang, id))
	if err != nil {
		return nil, err
	}
	charCollector.Wait()

	classJobCollector := collectors.BuildClassJobCollector(s.meta, s.profileSelectors, &charData)
	err = classJobCollector.Visit(fmt.Sprintf("https://%s.finalfantasyxiv.com/lodestone/character/%d/class_job/", s.lang, id))
	if err != nil {
		return nil, err
	}
	classJobCollector.Wait()

	return &charData, nil
}

// FetchCharacterMinions returns unlocked minion information for the provided Lodestone ID.
func (s *Scraper) FetchCharacterMinions(id uint32) ([]*models.Minion, error) {
	minionCollector := collectors.BuildMinionCollector(s.meta, s.profileSelectors)

	err := minionCollector.Visit(fmt.Sprintf("https://%s.finalfantasyxiv.com/lodestone/character/%d/minion/", s.lang, id))
	if err != nil {
		return nil, err
	}
	minionCollector.Wait()

	return nil, nil
}

// FetchCharacterMounts returns unlocked mount information for the provided Lodestone ID.
func (s *Scraper) FetchCharacterMounts(id uint32) ([]*models.Mount, error) {
	mountCollector := collectors.BuildMountCollector(s.meta, s.profileSelectors)

	err := mountCollector.Visit(fmt.Sprintf("https://%s.finalfantasyxiv.com/lodestone/character/%d/mount/", s.lang, id))
	if err != nil {
		return nil, err
	}
	mountCollector.Wait()

	return nil, nil
}

// FetchCharacterAchievements returns unlocked achievement information for the provided Lodestone ID.
func (s *Scraper) FetchCharacterAchievements(id uint32) chan *models.AchievementInfo {
	output := make(chan *models.AchievementInfo)

	go func() {
		achievementCollector := collectors.BuildAchievementCollector(s.meta, s.profileSelectors, output)
		err := achievementCollector.Visit(fmt.Sprintf("https://%s.finalfantasyxiv.com/lodestone/character/%d/achievement/", s.lang, id))
		if err != nil {
			output <- &models.AchievementInfo{
				Error: err,
			}
			close(output)
			return
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

	lsCollector := collectors.BuildLinkshellCollector(s.meta, s.linkshellSelectors, &ls)
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

	cwlsCollector := collectors.BuildCWLSCollector(s.meta, s.cwlsSelectors, &cwls)
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

	pvpTeamCollector := collectors.BuildPVPTeamCollector(s.meta, s.pvpTeamSelectors, &pvpTeam)
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

	fcCollector := collectors.BuildFreeCompanyCollector(s.meta, s.fcSelectors, &fc)
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
		fcMembersCollector := collectors.BuildFreeCompanyMembersCollector(s.meta, s.fcSelectors, output)

		err := fcMembersCollector.Visit(fmt.Sprintf("https://%s.finalfantasyxiv.com/lodestone/freecompany/%s/member/", s.lang, id))
		if err != nil {
			output <- &models.FreeCompanyMember{
				Error: err,
			}
			close(output)
			return
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
		searchCollector := collectors.BuildFreeCompanySearchCollector(s.meta, s.searchSelectors, output)

		err := searchCollector.Visit(uri)
		if err != nil {
			output <- &models.FreeCompanySearchResult{
				Error: err,
			}
			close(output)
			return
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
		searchCollector := collectors.BuildCharacterSearchCollector(s.meta, s.searchSelectors, output)

		err := searchCollector.Visit(uri)
		if err != nil {
			output <- &models.CharacterSearchResult{
				Error: err,
			}
			close(output)
			return
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
		searchCollector := collectors.BuildCWLSSearchCollector(s.meta, s.searchSelectors, output)

		err := searchCollector.Visit(uri)
		if err != nil {
			output <- &models.CWLSSearchResult{
				Error: err,
			}
			close(output)
			return
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
		searchCollector := collectors.BuildLinkshellSearchCollector(s.meta, s.searchSelectors, output)

		err := searchCollector.Visit(uri)
		if err != nil {
			output <- &models.LinkshellSearchResult{
				Error: err,
			}
			close(output)
			return
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
		searchCollector := collectors.BuildPVPTeamSearchCollector(s.meta, s.searchSelectors, output)

		err := searchCollector.Visit(uri)
		if err != nil {
			output <- &models.PVPTeamSearchResult{
				Error: err,
			}
			close(output)
			return
		}

		searchCollector.Wait()

		close(output)
	}()

	return output
}

// NewScraper creates a new instance of the Scraper.
func NewScraper(lang SiteLang) (*Scraper, error) {
	cwlsSelectors := selectors.LoadCWLSSelectors()
	lsSelectors := selectors.LoadLinkshellSelectors()
	profileSelectors := selectors.LoadProfileSelectors()
	pvpTeamSelectors := selectors.LoadPVPTeamSelectors()
	searchSelectors := selectors.LoadSearchSelectors()
	fcSelectors := selectors.LoadFreeCompanySelectors()

	metaBytes, _ := pack.Asset("meta.json")
	meta := models.Meta{}
	json.Unmarshal(metaBytes, &meta)

	return &Scraper{
		lang: lang,

		meta:               &meta,
		cwlsSelectors:      cwlsSelectors,
		fcSelectors:        fcSelectors,
		linkshellSelectors: lsSelectors,
		profileSelectors:   profileSelectors,
		pvpTeamSelectors:   pvpTeamSelectors,
		searchSelectors:    searchSelectors,
	}, nil
}
