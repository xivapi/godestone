package godestone

import (
	"fmt"

	"github.com/gocolly/colly/v2"
	"github.com/karashiiro/godestone/models"
)

// Scraper is the object through which interactions with The Lodestone are made.
type Scraper struct {
	charCollector   *colly.Collector
	minionCollector *colly.Collector
	mountCollector  *colly.Collector
	profSelectors   *profileSelectors
}

// FetchCharacter returns character information for the provided Lodestone ID.
func (s *Scraper) FetchCharacter(id uint32) (*models.Character, error) {
	charData := models.Character{}

	if s.charCollector == nil {
		s.charCollector = colly.NewCollector()
	}

	err := s.charCollector.Visit("https://na.finalfantasyxiv.com/lodestone/character/" + fmt.Sprint(id))
	if err != nil {
		return nil, err
	}

	return &charData, nil
}

// NewScraper creates a new instance of the Scraper.
func NewScraper() (*Scraper, error) {
	profSelectors, err := loadProfileSelectors()
	if err != nil {
		return nil, err
	}

	return &Scraper{
		profSelectors: profSelectors,
	}, nil
}
