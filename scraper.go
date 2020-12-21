package godestone

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/karashiiro/godestone/models"
	"github.com/karashiiro/godestone/pack"
	"github.com/karashiiro/godestone/selectors"
)

// Scraper is the object through which interactions with The Lodestone are made.
type Scraper struct {
	meta          *models.Meta
	profSelectors *selectors.ProfileSelectors
}

// FetchCharacter returns character information for the provided Lodestone ID.
func (s *Scraper) FetchCharacter(id uint32) (*models.Character, error) {
	now := time.Now()
	charData := models.Character{ID: id, ParseDate: &now}

	charCollector := s.makeCharCollector(&charData)
	err := charCollector.Visit("https://na.finalfantasyxiv.com/lodestone/character/" + fmt.Sprint(id))
	if err != nil {
		return nil, err
	}
	charCollector.Wait()

	classJobCollector := s.makeClassJobCollector(&charData)
	err = classJobCollector.Visit("https://na.finalfantasyxiv.com/lodestone/character/" + fmt.Sprint(id) + "/class_job/")
	if err != nil {
		return nil, err
	}
	classJobCollector.Wait()

	return &charData, nil
}

// FetchCharacterMinions returns unlocked minion information for the provided Lodestone ID.
func (s *Scraper) FetchCharacterMinions(id uint32) ([]*models.Minion, error) {
	minionCollector := s.makeMinionCollector()
	err := minionCollector.Visit("https://na.finalfantasyxiv.com/lodestone/character/" + fmt.Sprint(id) + "/minion/")
	if err != nil {
		return nil, err
	}
	minionCollector.Wait()

	return nil, nil
}

// FetchCharacterMounts returns unlocked mount information for the provided Lodestone ID.
func (s *Scraper) FetchCharacterMounts(id uint32) ([]*models.Mount, error) {
	mountCollector := s.makeMountCollector()
	err := mountCollector.Visit("https://na.finalfantasyxiv.com/lodestone/character/" + fmt.Sprint(id) + "/mount/")
	if err != nil {
		return nil, err
	}
	mountCollector.Wait()

	return nil, nil
}

// FetchCharacterAchievements returns unlocked achievement information for the provided Lodestone ID.
func (s *Scraper) FetchCharacterAchievements(id uint32) (*models.Achievements, error) {
	achievements := models.Achievements{}

	achievementCollector := s.makeAchievementCollector(&achievements)
	err := achievementCollector.Visit("https://na.finalfantasyxiv.com/lodestone/character/" + fmt.Sprint(id) + "/achievement/")
	if err != nil {
		return nil, err
	}
	achievementCollector.Wait()

	return &achievements, nil
}

// NewScraper creates a new instance of the Scraper.
func NewScraper() (*Scraper, error) {
	profSelectors, err := selectors.LoadProfileSelectors()
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
		meta:          &meta,
		profSelectors: profSelectors,
	}, nil
}
