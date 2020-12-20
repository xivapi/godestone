package godestone

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

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

		s.charCollector.OnHTML(s.profSelectors.Character["AVATAR"].(string), func(e *colly.HTMLElement) {
			charData.Avatar = e.Attr("src")
		})

		s.charCollector.OnHTML(s.profSelectors.Character["BIO"].(string), func(e *colly.HTMLElement) {
			charData.Bio = e.Text
		})

		s.charCollector.OnHTML(s.profSelectors.Character["NAME"].(string), func(e *colly.HTMLElement) {
			charData.Name = e.Text
		})

		s.charCollector.OnHTML(s.profSelectors.Character["NAMEDAY"].(string), func(e *colly.HTMLElement) {
			charData.Nameday = e.Text
		})

		s.charCollector.OnHTML(s.profSelectors.Character["PORTRAIT"].(string), func(e *colly.HTMLElement) {
			charData.Portrait = e.Attr("src")
		})

		pvpTeamIDRegex := regexp.MustCompile("/lodestone/pvpteam/(?P<ID>.+)/")
		s.charCollector.OnHTML(s.profSelectors.Character["PVP_TEAM"].(map[string](interface{}))["NAME"].(string), func(e *colly.HTMLElement) {
			matches := pvpTeamIDRegex.FindStringSubmatch(e.Attr("href"))
			if matches != nil {
				charData.PvPTeamID = matches[1]
			}
		})

		s.charCollector.OnHTML(s.profSelectors.Character["SERVER"].(string), func(e *colly.HTMLElement) {
			server := e.Text
			serverSplit := strings.Split(server, "(")
			world := serverSplit[0][0 : len(serverSplit[0])-2]
			dc := serverSplit[1][0 : len(serverSplit[1])-1]

			charData.DC = dc
			charData.Server = world
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
