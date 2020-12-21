package godestone

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/karashiiro/godestone/data/grandcompany"
	"github.com/karashiiro/godestone/data/race"
	"github.com/karashiiro/godestone/data/tribe"
	"github.com/karashiiro/godestone/models"
	"github.com/karashiiro/godestone/pack"
	"github.com/karashiiro/godestone/selectors"
)

// Scraper is the object through which interactions with The Lodestone are made.
type Scraper struct {
	meta             *models.Meta
	profileSelectors *selectors.ProfileSelectors
	searchSelectors  *selectors.SearchSelectors
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

// SearchCharacterOptions defines extra search information that can help to narrow down a search.
type SearchCharacterOptions struct {
	Name         string
	World        string
	DC           string
	Lang         Lang
	GrandCompany grandcompany.GrandCompany
	Race         race.Race
	Tribe        tribe.Tribe
	Order        SearchOrder
}

// SearchCharacters returns a channel of searchable characters.
func (s *Scraper) SearchCharacters(opts SearchCharacterOptions) chan *models.CharacterSearchResult {
	output := make(chan *models.CharacterSearchResult)

	uriFormat := "https://na.finalfantasyxiv.com/lodestone/character/?q=%s&worldname=%s&classjob=%s&order=%d"

	name := strings.Replace(opts.Name, " ", "%20", -1)

	worldDC := opts.DC
	if len(opts.World) != 0 {
		worldDC = opts.World
	} else {
		// DCs have the _dc_ prefix attached to them
		if len(worldDC) != 0 && !strings.HasPrefix(worldDC, "_dc_") {
			worldDC = "_dc_" + worldDC
		}
	}

	if opts.Lang == None || opts.Lang&JA != 0 {
		uriFormat += "&blog_lang=ja"
	}
	if opts.Lang == None || opts.Lang&EN != 0 {
		uriFormat += "&blog_lang=en"
	}
	if opts.Lang == None || opts.Lang&DE != 0 {
		uriFormat += "&blog_lang=de"
	}
	if opts.Lang == None || opts.Lang&FR != 0 {
		uriFormat += "&blog_lang=fr"
	}

	if opts.Tribe != tribe.None || opts.Race != race.None {
		raceTribe := ""
		if opts.Tribe != tribe.None {
			raceTribe = fmt.Sprintf("tribe_%d", opts.Tribe)
		} else if opts.Race != race.None {
			raceTribe = fmt.Sprintf("race_%d", opts.Race)
		}
		uriFormat += fmt.Sprintf("&race_tribe=%s", raceTribe)
	}

	if opts.GrandCompany != grandcompany.None {
		uriFormat += fmt.Sprintf("&gcid=%d", opts.GrandCompany)
	}

	builtURI := fmt.Sprintf(uriFormat, name, worldDC, "", opts.Order)

	go func() {
		searchCollector := s.makeCharacterSearchCollector(output)

		err := searchCollector.Visit(builtURI)
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

// NewScraper creates a new instance of the Scraper.
func NewScraper() (*Scraper, error) {
	profileSelectors, err := selectors.LoadProfileSelectors()
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
		searchSelectors:  searchSelectors,
	}, nil
}
