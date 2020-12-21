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
	meta          map[string]string
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

	return &charData, nil
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
	meta := make(map[string]string)
	json.Unmarshal(metaBytes, &meta)

	return &Scraper{
		meta:          meta,
		profSelectors: profSelectors,
	}, nil
}
