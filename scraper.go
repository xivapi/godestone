package godestone

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/karashiiro/godestone/collectors"

	"github.com/karashiiro/godestone/models"
	"github.com/karashiiro/godestone/pack"
	"github.com/karashiiro/godestone/selectors"
)

// Scraper is the object through which interactions with The Lodestone are made.
type Scraper struct {
	meta             *models.Meta
	profileSelectors *selectors.ProfileSelectors
	pvpTeamSelectors *selectors.PVPTeamSelectors
	searchSelectors  *selectors.SearchSelectors
}

// FetchCharacter returns character information for the provided Lodestone ID.
func (s *Scraper) FetchCharacter(id uint32) (*models.Character, error) {
	now := time.Now()
	charData := models.Character{ID: id, ParseDate: now}

	charCollector := collectors.BuildCharacterCollector(s.meta, s.profileSelectors, &charData)
	err := charCollector.Visit("https://na.finalfantasyxiv.com/lodestone/character/" + fmt.Sprint(id))
	if err != nil {
		return nil, err
	}
	charCollector.Wait()

	classJobCollector := collectors.BuildClassJobCollector(s.meta, s.profileSelectors, &charData)
	err = classJobCollector.Visit("https://na.finalfantasyxiv.com/lodestone/character/" + fmt.Sprint(id) + "/class_job/")
	if err != nil {
		return nil, err
	}
	classJobCollector.Wait()

	return &charData, nil
}

// FetchCharacterMinions returns unlocked minion information for the provided Lodestone ID.
func (s *Scraper) FetchCharacterMinions(id uint32) ([]*models.Minion, error) {
	minionCollector := collectors.BuildMinionCollector(s.meta, s.profileSelectors)

	err := minionCollector.Visit("https://na.finalfantasyxiv.com/lodestone/character/" + fmt.Sprint(id) + "/minion/")
	if err != nil {
		return nil, err
	}
	minionCollector.Wait()

	return nil, nil
}

// FetchCharacterMounts returns unlocked mount information for the provided Lodestone ID.
func (s *Scraper) FetchCharacterMounts(id uint32) ([]*models.Mount, error) {
	mountCollector := collectors.BuildMountCollector(s.meta, s.profileSelectors)

	err := mountCollector.Visit("https://na.finalfantasyxiv.com/lodestone/character/" + fmt.Sprint(id) + "/mount/")
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
		err := achievementCollector.Visit("https://na.finalfantasyxiv.com/lodestone/character/" + fmt.Sprint(id) + "/achievement/")
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

// FetchPVPTeam returns PVP team information for the provided PVP team ID.
func (s *Scraper) FetchPVPTeam(id string) (*models.PVPTeam, error) {
	now := time.Now()
	pvpTeam := models.PVPTeam{ID: id, ParseDate: now}

	pvpTeamCollector := collectors.BuildPVPTeamCollector(s.meta, s.pvpTeamSelectors, &pvpTeam)
	err := pvpTeamCollector.Visit("https://na.finalfantasyxiv.com/lodestone/pvpteam/" + fmt.Sprint(id))
	if err != nil {
		return nil, err
	}
	pvpTeamCollector.Wait()

	return &pvpTeam, nil
}

// SearchCharacters returns a channel of searchable characters.
func (s *Scraper) SearchCharacters(opts SearchCharacterOptions) chan *models.CharacterSearchResult {
	output := make(chan *models.CharacterSearchResult)

	uri := opts.buildURI()
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
func (s *Scraper) SearchCWLS(opts SearchCWLSOptions) chan *models.CWLSSearchResult {
	output := make(chan *models.CWLSSearchResult)

	uri := opts.buildURI()
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
func (s *Scraper) SearchLinkshells(opts SearchLinkshellOptions) chan *models.LinkshellSearchResult {
	output := make(chan *models.LinkshellSearchResult)

	uri := opts.buildURI()
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
func (s *Scraper) SearchPVPTeams(opts SearchPVPTeamOptions) chan *models.PVPTeamSearchResult {
	output := make(chan *models.PVPTeamSearchResult)

	uri := opts.buildURI()
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
func NewScraper() (*Scraper, error) {
	profileSelectors, err := selectors.LoadProfileSelectors()
	if err != nil {
		return nil, err
	}

	pvpTeamSelectors, err := selectors.LoadPVPTeamSelectors()
	if err != nil {
		return nil, err
	}

	searchSelectors, err := selectors.LoadSearchSelectors()
	if err != nil {
		return nil, err
	}

	metaBytes, err := pack.Asset("meta.json")
	if err != nil {
		return nil, err
	}
	meta := models.Meta{}
	json.Unmarshal(metaBytes, &meta)

	return &Scraper{
		meta:             &meta,
		profileSelectors: profileSelectors,
		pvpTeamSelectors: pvpTeamSelectors,
		searchSelectors:  searchSelectors,
	}, nil
}
