package collectors

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/gocolly/colly/v2"
	"github.com/karashiiro/godestone/models"
	"github.com/karashiiro/godestone/pack/exports"
	"github.com/karashiiro/godestone/selectors"
	lookups "github.com/karashiiro/godestone/table-lookups"
)

var nonDigits = regexp.MustCompile("[^\\d]")

// BuildClassJobCollector builds the collector used for processing the page.
func BuildClassJobCollector(meta *models.Meta, profSelectors *selectors.ProfileSelectors, classJobTable *exports.ClassJobTable, charData *models.Character) *colly.Collector {
	c := colly.NewCollector()
	c.UserAgent = meta.UserAgentDesktop
	c.IgnoreRobotsTxt = true

	classJobSelectors := profSelectors.ClassJob
	charData.ClassJobs = make([]*models.ClassJob, 0)

	cjRefs := []*selectors.OneClassJobSelectors{
		&classJobSelectors.Paladin,
		&classJobSelectors.Warrior,
		&classJobSelectors.DarkKnight,
		&classJobSelectors.Gunbreaker,
		&classJobSelectors.Monk,
		&classJobSelectors.Dragoon,
		&classJobSelectors.Ninja,
		&classJobSelectors.Samurai,
		&classJobSelectors.WhiteMage,
		&classJobSelectors.Scholar,
		&classJobSelectors.Astrologian,
		&classJobSelectors.Bard,
		&classJobSelectors.Machinist,
		&classJobSelectors.Dancer,
		&classJobSelectors.BlackMage,
		&classJobSelectors.Summoner,
		&classJobSelectors.RedMage,
		&classJobSelectors.BlueMage,
		&classJobSelectors.Carpenter,
		&classJobSelectors.Blacksmith,
		&classJobSelectors.Armorer,
		&classJobSelectors.Goldsmith,
		&classJobSelectors.Leatherworker,
		&classJobSelectors.Weaver,
		&classJobSelectors.Alchemist,
		&classJobSelectors.Culinarian,
		&classJobSelectors.Miner,
		&classJobSelectors.Botanist,
		&classJobSelectors.Fisher,
	}

	for _, ref := range cjRefs {
		curRef := ref
		curCj := models.ClassJob{}

		c.OnHTML(curRef.Exp.Selector, func(e *colly.HTMLElement) {
			expStrs := curRef.Exp.Parse(e)
			expStrs[0] = nonDigits.ReplaceAllString(expStrs[0], "")
			expStrs[1] = nonDigits.ReplaceAllString(expStrs[1], "")

			curExp, err := strconv.ParseUint(expStrs[0], 10, 32)
			if err == nil {
				curCj.ExpLevel = uint32(curExp)
			}

			maxExp, err := strconv.ParseUint(expStrs[1], 10, 32)
			if err == nil {
				curCj.ExpLevelMax = uint32(maxExp)
			}

			curCj.ExpLevelTogo = curCj.ExpLevelMax - curCj.ExpLevel
		})

		c.OnHTML(curRef.Level.Selector, func(e *colly.HTMLElement) {
			levelStr := curRef.Level.Parse(e)[0]
			level, err := strconv.ParseUint(levelStr, 10, 8)
			if err == nil {
				curCj.Level = uint8(level)
			}
		})

		c.OnHTML(curRef.UnlockState.Selector, func(e *colly.HTMLElement) {
			curCj.UnlockedState.Name = curRef.UnlockState.Parse(e)[0]

			// This bit is kinda wack, should fix it
			curCj.Name = e.Attr("data-tooltip")
			names := strings.Split(curCj.Name, " / ")
			names[0] = strings.TrimRight(strings.Split(names[0], "(")[0], " ")
			names[0] = strings.TrimRight(strings.Split(names[0], "[")[0], " ")

			curCj.IsSpecialized = strings.Contains(e.Attr("class"), "meister")

			jobInfo := lookups.ClassJobTableLookup(classJobTable, names[0])
			curCj.JobID = uint8(jobInfo.Id())

			if len(names) > 1 {
				classInfo := lookups.ClassJobTableLookup(classJobTable, names[1])
				curCj.ClassID = uint8(classInfo.Id())
			} else {
				curCj.ClassID = uint8(curCj.JobID)
			}

			cjInfo := lookups.ClassJobTableLookup(classJobTable, curCj.UnlockedState.Name)
			curCj.UnlockedState.ID = uint8(cjInfo.Id())
		})

		charData.ClassJobs = append(charData.ClassJobs, &curCj)
	}

	cjb := &models.ClassJobBozja{}
	c.OnHTML(classJobSelectors.Bozja.Level.Selector, func(e *colly.HTMLElement) {
		levelStr := classJobSelectors.Bozja.Level.Parse(e)[0]
		level, err := strconv.ParseUint(levelStr, 10, 8)
		if err == nil {
			cjb.Level = uint8(level)
		}
	})

	c.OnHTML(classJobSelectors.Bozja.Mettle.Selector, func(e *colly.HTMLElement) {
		mettleStr := classJobSelectors.Bozja.Mettle.Parse(e)[0]
		mettleStr = nonDigits.ReplaceAllString(mettleStr, "")
		mettle, err := strconv.ParseUint(mettleStr, 10, 32)
		if err == nil {
			cjb.Mettle = uint32(mettle)
		}
	})

	c.OnHTML(classJobSelectors.Bozja.Name.Selector, func(e *colly.HTMLElement) {
		cjb.Name = classJobSelectors.Bozja.Name.Parse(e)[0]
	})
	charData.ClassJobBozjan = cjb

	cje := &models.ClassJobEureka{}
	c.OnHTML(classJobSelectors.Eureka.Level.Selector, func(e *colly.HTMLElement) {
		levelStr := classJobSelectors.Eureka.Level.Parse(e)[0]
		level, err := strconv.ParseUint(levelStr, 10, 8)
		if err == nil {
			cje.Level = uint8(level)
		}
	})

	c.OnHTML(classJobSelectors.Eureka.Exp.Selector, func(e *colly.HTMLElement) {
		expStrs := classJobSelectors.Eureka.Exp.Parse(e)
		expStrs[0] = nonDigits.ReplaceAllString(expStrs[0], "")
		expStrs[1] = nonDigits.ReplaceAllString(expStrs[1], "")

		curExp, err := strconv.ParseUint(expStrs[0], 10, 32)
		if err == nil {
			cje.ExpLevel = uint32(curExp)
		}

		maxExp, err := strconv.ParseUint(expStrs[1], 10, 32)
		if err == nil {
			cje.ExpLevelMax = uint32(maxExp)
		}

		cje.ExpLevelTogo = cje.ExpLevelMax - cje.ExpLevel
	})

	c.OnHTML(classJobSelectors.Eureka.Name.Selector, func(e *colly.HTMLElement) {
		cje.Name = classJobSelectors.Eureka.Name.Parse(e)[0]
	})
	charData.ClassJobElemental = cje

	return c
}
