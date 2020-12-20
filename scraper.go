package godestone

import (
	"encoding/json"
	"fmt"

	"github.com/gocolly/colly/v2"
	"github.com/karashiiro/godestone/models"
	"github.com/karashiiro/godestone/pack"
)

// Scraper is the object through which interactions with The Lodestone are made.
type Scraper struct {
	charCollector   *colly.Collector
	minionCollector *colly.Collector
	mountCollector  *colly.Collector

	meta          map[string]string
	profSelectors *profileSelectors
}

// FetchCharacter returns character information for the provided Lodestone ID.
func (s *Scraper) FetchCharacter(id uint32) (*models.Character, error) {
	charData := models.Character{}

	if s.charCollector == nil {
		s.charCollector = colly.NewCollector()
		s.charCollector.UserAgent = s.meta["userAgentDesktop"]
		s.charCollector.IgnoreRobotsTxt = true

		s.charCollector.OnHTML(s.profSelectors.Character["AVATAR"], func(e *colly.HTMLElement) {
			charData.Avatar = e.Attr("src")
		})
	}

	err := s.charCollector.Visit("https://na.finalfantasyxiv.com/lodestone/character/" + fmt.Sprint(id))
	if err != nil {
		return nil, err
	}

	s.charCollector.Wait()

	return &charData, nil
}

// NewScraper creates a new instance of the Scraper.
func NewScraper() (*Scraper, error) {
	profSelectors, err := loadProfileSelectors()
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
