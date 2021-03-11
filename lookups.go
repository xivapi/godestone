package godestone

import (
	"regexp"
	"strings"

	"github.com/xivapi/godestone/pack/exports"
)

func listContains(test string, list ...string) bool {
	for _, listItem := range list {
		itemLower := removeBracketedPhrases(strings.ToLower(listItem))
		if strings.Contains(itemLower, test) {
			return true
		}
	}

	return false
}

var bracketed = regexp.MustCompile("\\[.*\\]")

func removeBracketedPhrases(input string) string {
	return bracketed.ReplaceAllString(input, "")
}

func (s *Scraper) achievementTableLookup(name string) *exports.Achievement {
	nameLower := strings.ToLower(name)

	nAchievements := s.getAchievementTable().AchievementsLength()
	for i := 0; i < nAchievements; i++ {
		achievement := exports.Achievement{}
		s.getAchievementTable().Achievements(&achievement, i)

		nameEn := string(achievement.NameEn())
		nameDe := string(achievement.NameDe())
		nameFr := string(achievement.NameFr())
		nameJa := string(achievement.NameJa())

		if listContains(
			nameLower,
			nameEn,
			nameDe,
			nameFr,
			nameJa,
		) {
			return &achievement
		}
	}

	return nil
}

func (s *Scraper) classJobTableLookup(name string) *exports.ClassJob {
	nameLower := strings.ToLower(name)

	nClassJobs := s.getClassJobTable().ClassJobsLength()
	for i := 0; i < nClassJobs; i++ {
		cj := exports.ClassJob{}
		s.getClassJobTable().ClassJobs(&cj, i)

		nameEn := string(cj.NameEn())
		nameDe := string(cj.NameDe())
		nameFr := string(cj.NameFr())
		nameJa := string(cj.NameJa())

		if listContains(
			nameLower,
			nameEn,
			nameDe,
			nameFr,
			nameJa,
		) {
			return &cj
		}
	}

	return nil
}

func (s *Scraper) deityTableLookup(name string) *exports.Deity {
	nameLower := strings.ToLower(name)

	nDeities := s.getDeityTable().DeitiesLength()
	for i := 0; i < nDeities; i++ {
		deity := exports.Deity{}
		s.getDeityTable().Deities(&deity, i)

		nameEn := string(deity.NameEn())
		nameDe := string(deity.NameDe())
		nameFr := string(deity.NameFr())
		nameJa := string(deity.NameJa())

		if listContains(
			nameLower,
			nameEn,
			nameDe,
			nameFr,
			nameJa,
		) {
			return &deity
		}
	}

	return nil
}

func (s *Scraper) grandCompanyTableLookup(name string) *exports.GrandCompany {
	nameLower := strings.ToLower(name)

	nGCs := s.getGrandCompanyTable().GrandCompaniesLength()
	for i := 0; i < nGCs; i++ {
		gc := exports.GrandCompany{}
		s.getGrandCompanyTable().GrandCompanies(&gc, i)

		nameEn := string(gc.NameEn())
		nameJa := string(gc.NameJa())
		nameDe := string(gc.NameDe())
		nameFr := string(gc.NameFr())

		if listContains(
			nameLower,
			nameEn,
			nameDe,
			nameFr,
			nameJa,
		) {
			return &gc
		}
	}

	return nil
}

func (s *Scraper) itemTableLookup(name string) *exports.Item {
	nameLower := strings.ToLower(name)

	nItems := s.getItemTable().ItemsLength()
	for i := 0; i < nItems; i++ {
		item := exports.Item{}
		s.getItemTable().Items(&item, i)

		nameEn := string(item.NameEn())
		nameDe := string(item.NameDe())
		nameFr := string(item.NameFr())
		nameJa := string(item.NameJa())

		if listContains(
			nameLower,
			nameEn,
			nameDe,
			nameFr,
			nameJa,
		) {
			return &item
		}
	}

	return nil
}

func (s *Scraper) minionTableLookup(name string) *exports.Minion {
	nameLower := strings.ToLower(name)

	// Thanks, German
	// If anyone knows how to properly handle this, a PR would be more than welcome.
	if string(s.lang) == "de" {
		nameLower = strings.Replace(nameLower, "blaublütiger ", "baby-", 1)
		nameLower = strings.Replace(nameLower, "es ", " ", 1)
		if !strings.HasPrefix(nameLower, "seite") {
			nameLower = strings.Replace(nameLower, "e ", " ", 1)
		}
		nameLower = strings.Replace(nameLower, "er ", " ", 1)
		nameLower = strings.Replace(nameLower, " d ", " der ", 1)
		if strings.Contains(nameLower, "chocobo-küken") {
			parts := strings.Split(nameLower, " ")
			nameLower = parts[len(parts)-1]
		}
	}

	nMinions := s.getMinionTable().MinionsLength()
	for i := 0; i < nMinions; i++ {
		minion := exports.Minion{}
		s.getMinionTable().Minions(&minion, i)

		nameEn := string(minion.NameEn())
		nameDe := string(minion.NameDe())
		nameFr := string(minion.NameFr())
		nameJa := string(minion.NameJa())

		if listContains(
			nameLower,
			nameEn,
			nameDe,
			nameFr,
			nameJa,
		) {
			return &minion
		}
	}

	return nil
}

func (s *Scraper) mountTableLookup(name string) *exports.Mount {
	nameLower := strings.ToLower(name)

	nMounts := s.getMountTable().MountsLength()
	for i := 0; i < nMounts; i++ {
		mount := exports.Mount{}
		s.getMountTable().Mounts(&mount, i)

		nameEn := string(mount.NameEn())
		nameDe := string(mount.NameDe())
		nameFr := string(mount.NameFr())
		nameJa := string(mount.NameJa())

		if listContains(
			nameLower,
			nameEn,
			nameDe,
			nameFr,
			nameJa,
		) {
			return &mount
		}
	}

	return nil
}

func (s *Scraper) raceTableLookup(name string) *exports.Race {
	nameLower := strings.ToLower(name)

	nRaces := s.getRaceTable().RacesLength()
	for i := 0; i < nRaces; i++ {
		race := exports.Race{}
		s.getRaceTable().Races(&race, i)

		nameMasculineEn := string(race.NameMasculineEn())
		nameMasculineDe := string(race.NameMasculineDe())
		nameMasculineFr := string(race.NameMasculineFr())
		nameMasculineJa := string(race.NameMasculineJa())
		nameFeminineEn := string(race.NameFeminineEn())
		nameFeminineDe := string(race.NameFeminineDe())
		nameFeminineFr := string(race.NameFeminineFr())
		nameFeminineJa := string(race.NameFeminineJa())

		if listContains(
			nameLower,
			nameMasculineEn,
			nameMasculineDe,
			nameMasculineFr,
			nameMasculineJa,
			nameFeminineEn,
			nameFeminineDe,
			nameFeminineFr,
			nameFeminineJa,
		) {
			return &race
		}
	}

	return nil
}

func (s *Scraper) reputationTableLookup(name string) *exports.Reputation {
	nameLower := strings.ToLower(name)

	nReputations := s.getReputationTable().ReputationsLength()
	for i := 0; i < nReputations; i++ {
		reputation := exports.Reputation{}
		s.getReputationTable().Reputations(&reputation, i)

		nameEn := string(reputation.NameEn())
		nameDe := string(reputation.NameDe())
		nameFr := string(reputation.NameFr())
		nameJa := string(reputation.NameJa())

		if listContains(
			nameLower,
			nameEn,
			nameDe,
			nameFr,
			nameJa,
		) {
			return &reputation
		}
	}

	return nil
}

func (s *Scraper) titleTableLookup(name string) *exports.Title {
	nameLower := strings.ToLower(name)

	nTitles := s.getTitleTable().TitlesLength()
	for i := 0; i < nTitles; i++ {
		title := exports.Title{}
		s.getTitleTable().Titles(&title, i)

		nameMasculineEn := string(title.NameMasculineEn())
		nameMasculineDe := string(title.NameMasculineDe())
		nameMasculineFr := string(title.NameMasculineFr())
		nameMasculineJa := string(title.NameMasculineJa())
		nameFeminineEn := string(title.NameFeminineEn())
		nameFeminineDe := string(title.NameFeminineDe())
		nameFeminineFr := string(title.NameFeminineFr())
		nameFeminineJa := string(title.NameFeminineJa())

		if listContains(
			nameLower,
			nameMasculineEn,
			nameMasculineDe,
			nameMasculineFr,
			nameMasculineJa,
			nameFeminineEn,
			nameFeminineDe,
			nameFeminineFr,
			nameFeminineJa,
		) {
			return &title
		}
	}

	return nil
}

func (s *Scraper) townTableLookup(name string) *exports.Town {
	nameLower := strings.ToLower(name)

	nTowns := s.getTownTable().TownsLength()
	for i := 0; i < nTowns; i++ {
		town := exports.Town{}
		s.getTownTable().Towns(&town, i)

		nameEn := string(town.NameEn())
		nameDe := string(town.NameDe())
		nameFr := string(town.NameFr())
		nameJa := string(town.NameJa())

		if listContains(
			nameLower,
			nameEn,
			nameDe,
			nameFr,
			nameJa,
		) {
			return &town
		}
	}

	return nil
}

func (s *Scraper) tribeTableLookup(name string) *exports.Tribe {
	nameLower := strings.ToLower(name)

	nTribes := s.getTribeTable().TribesLength()
	for i := 0; i < nTribes; i++ {
		tribe := exports.Tribe{}
		s.getTribeTable().Tribes(&tribe, i)

		nameMasculineEn := string(tribe.NameMasculineEn())
		nameMasculineDe := string(tribe.NameMasculineDe())
		nameMasculineFr := string(tribe.NameMasculineFr())
		nameMasculineJa := string(tribe.NameMasculineJa())
		nameFeminineEn := string(tribe.NameFeminineEn())
		nameFeminineDe := string(tribe.NameFeminineDe())
		nameFeminineFr := string(tribe.NameFeminineFr())
		nameFeminineJa := string(tribe.NameFeminineJa())

		if listContains(
			nameLower,
			nameMasculineEn,
			nameMasculineDe,
			nameMasculineFr,
			nameMasculineJa,
			nameFeminineEn,
			nameFeminineDe,
			nameFeminineFr,
			nameFeminineJa,
		) {
			return &tribe
		}
	}

	return nil
}
