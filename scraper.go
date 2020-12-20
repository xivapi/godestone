package godestone

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/karashiiro/godestone/data/deity"
	"github.com/karashiiro/godestone/data/gcrank"
	"github.com/karashiiro/godestone/data/gender"
	"github.com/karashiiro/godestone/data/grandcompany"
	"github.com/karashiiro/godestone/data/race"
	"github.com/karashiiro/godestone/data/town"
	"github.com/karashiiro/godestone/data/tribe"
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
	now := time.Now()
	charData := models.Character{ID: id, ParseDate: &now}

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

		fcIDRegex := regexp.MustCompile("/lodestone/freecompany/(?P<ID>.+)/")
		s.charCollector.OnHTML(s.profSelectors.Character["FREE_COMPANY"].(map[string](interface{}))["NAME"].(string), func(e *colly.HTMLElement) {
			matches := fcIDRegex.FindStringSubmatch(e.Attr("href"))
			if matches != nil {
				/*
					This could be parsed to a uint64, but I don't know what SE's theoretical cap on Free Companies is and I'd
					rather this not break in a decade. It's harmless to keep it as a string, anyways, since it needs to be
					onverted to one to do a Lodestone lookup with it and anyone who wants it as a uint64 can just convert it themselves.
				*/
				charData.FreeCompanyID = matches[1]
				charData.FreeCompanyName = e.Text
			}
		})

		s.charCollector.OnHTML(s.profSelectors.Character["GRAND_COMPANY"].(string), func(e *colly.HTMLElement) {
			gcRawInfo := strings.Split(e.Text, "/")
			gcName := gcRawInfo[0][0 : len(gcRawInfo[0])-1]
			gcRankNameParts := strings.Split(gcRawInfo[1][1:], " ")
			gcRank := gcRankNameParts[len(gcRankNameParts)-1]

			gcID := grandcompany.Parse(gcName)
			gcRankID := gcrank.Parse(gcRank)

			gc := models.GrandCompanyInfo{NameID: gcID, RankID: gcRankID}
			charData.GrandCompany = &gc
		})

		s.charCollector.OnHTML(s.profSelectors.Character["GUARDIAN_DEITY"].(string), func(e *colly.HTMLElement) {
			charData.GuardianDeity = deity.Parse(e.Text)
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

		raceClanGenderRegex := regexp.MustCompile("(?P<Race>.*)<br\\/>(?P<Tribe>.*) \\/ (?P<Gender>.)")
		s.charCollector.OnHTML(s.profSelectors.Character["RACE_CLAN_GENDER"].(string), func(e *colly.HTMLElement) {
			rawText, err := e.DOM.Html()
			if err != nil {
				return
			}

			matches := raceClanGenderRegex.FindStringSubmatch(rawText)
			if matches != nil {
				charData.Race = race.Parse(matches[1])
				charData.Tribe = tribe.Parse(matches[2])
				charData.Gender = gender.Parse(matches[3])
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

		s.charCollector.OnHTML(s.profSelectors.Character["TITLE"].(string), func(e *colly.HTMLElement) {
			// TODO
			charData.Title = 0
			charData.TitleTop = false
		})

		s.charCollector.OnHTML(s.profSelectors.Character["TOWN"].(string), func(e *colly.HTMLElement) {
			charData.Town = town.Parse(e.Text)
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
