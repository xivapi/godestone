package godestone

import (
	"fmt"
	"testing"
	"time"

	"github.com/karashiiro/godestone/data/grandcompany"
	"github.com/karashiiro/godestone/search"
)

var langCodes []SiteLang = []SiteLang{EN, JA, FR, DE, SiteLang("zh")}

func failIfStringEmpty(t *testing.T, label string, input string) {
	if input == "" {
		t.Errorf(fmt.Sprintf("%s is empty; expected non-empty string", label))
	}
}

func failIfNumberZero(t *testing.T, label string, input int64) {
	if input == 0 {
		t.Errorf(fmt.Sprintf("%s is zero; expected nonzero value", label))
	}
}

func failIfOlderThanGameRelease(t *testing.T, label string, input time.Time) {
	release, _ := time.Parse("2006-Jan-02", "2013-Aug-16") // Early access began on the 24th, but there are multiple timestamps for the 16th. Wikipedia probably left something out.
	inputStr, _ := input.MarshalText()
	if input.Before(release) {
		t.Errorf(fmt.Sprintf("%s is older than the game's release date; got %s", label, string(inputStr)))
	}
}

func failIfGCInvalid(t *testing.T, label string, input grandcompany.GrandCompany) {
	if grandcompany.Parse(string(input)) != grandcompany.None {
		t.Errorf(fmt.Sprintf("%s is not a valid Grand Company; got %s", label, string(input)))
	}
}

var linkshellIds []string = []string{"20547673299961415", "19703248369746483", "10414574138338845"}

func TestFetchLinkshell(t *testing.T) {
	for _, lang := range langCodes {
		s := NewScraper(lang)

		t.Run("SiteLang: "+string(lang), func(t *testing.T) {
			for _, id := range linkshellIds {
				t.Run("Linkshell ID "+id, func(t *testing.T) {
					ls, err := s.FetchLinkshell(id)
					if err != nil {
						if lang == SiteLang("zh") { // A-OK, there is no Chinese website
							return
						}

						t.Errorf(err.Error())
					}

					failIfStringEmpty(t, "Linkshell name", ls.Name)
					failIfStringEmpty(t, "Linkshell ID", ls.ID)

					if len(ls.Members) == 0 {
						t.Errorf("Linkshell has no members")
					}

					for _, member := range ls.Members {
						failIfStringEmpty(t, "Member avatar", member.Avatar)
						failIfNumberZero(t, "Member ID", int64(member.ID))
						failIfStringEmpty(t, "Member name", member.Name)
						failIfStringEmpty(t, "Member world", member.World)
						failIfStringEmpty(t, "Member DC", member.DC)
					}
				})
			}
		})
	}
}

var cwlsIds []string = []string{
	"4b8af89f50a062b4b15650ecf6583f7ac9ad8065",
	"4e7baf2e534e3fcd13edf24f554ddeb8b9efa1b5",
	"3b417d2c5390d9ebf62d35bd63f67fe26eb3d828",
}

func TestFetchCWLS(t *testing.T) {
	for _, lang := range langCodes {
		s := NewScraper(lang)

		t.Run("SiteLang: "+string(lang), func(t *testing.T) {
			for _, id := range cwlsIds {
				t.Run("CWLS ID "+id, func(t *testing.T) {
					cwls, err := s.FetchCWLS(id)
					if err != nil {
						if lang == SiteLang("zh") {
							return
						}

						t.Errorf(err.Error())
					}

					failIfStringEmpty(t, "CWLS name", cwls.Name)
					failIfStringEmpty(t, "CWLS ID", cwls.ID)
					failIfStringEmpty(t, "CWLS DC", cwls.DC)

					if len(cwls.Members) == 0 {
						t.Errorf("CWLS has no members")
					}

					for _, member := range cwls.Members {
						failIfStringEmpty(t, "Member avatar", member.Avatar)
						failIfNumberZero(t, "Member ID", int64(member.ID))
						failIfStringEmpty(t, "Member name", member.Name)
						failIfStringEmpty(t, "Member world", member.World)
						failIfStringEmpty(t, "Member DC", member.DC)
					}
				})
			}
		})
	}
}

var pvpTeamIds []string = []string{
	"253c62269c624bc115902cea98e84fe082b79f85",
	"a9b97f78cd9a59a6c71adb6d35ca8f902faf12d6",
	"bbe7823327192ab12ad5b8215f5d07f1b8edabed",
}

func TestFetchPVPTeam(t *testing.T) {
	for _, lang := range langCodes {
		s := NewScraper(lang)

		t.Run("SiteLang: "+string(lang), func(t *testing.T) {
			for _, id := range pvpTeamIds {
				t.Run("PVP team ID "+id, func(t *testing.T) {
					pvpTeam, err := s.FetchPVPTeam(id)
					if err != nil {
						if lang == SiteLang("zh") {
							return
						}

						t.Errorf(err.Error())
					}

					failIfStringEmpty(t, "PVP team name", pvpTeam.Name)
					failIfStringEmpty(t, "PVP team ID", pvpTeam.ID)
					failIfStringEmpty(t, "PVP team DC", pvpTeam.DC)

					failIfOlderThanGameRelease(t, "PVP team formation", pvpTeam.Formed)

					if len(pvpTeam.Members) == 0 {
						t.Errorf("PVP team has no members")
					}

					for _, member := range pvpTeam.Members {
						failIfStringEmpty(t, "Member avatar", member.Avatar)
						failIfNumberZero(t, "Member ID", int64(member.ID))
						failIfStringEmpty(t, "Member name", member.Name)
						failIfStringEmpty(t, "Member world", member.World)
						failIfStringEmpty(t, "Member DC", member.DC)
					}
				})
			}
		})
	}
}

var fcIds []string = []string{
	"9231816286156096656",
	"9230268173784187532",
	"9232660711086230486",
}

func TestFetchFreeCompany(t *testing.T) {
	for _, lang := range langCodes {
		s := NewScraper(lang)

		t.Run("SiteLang: "+string(lang), func(t *testing.T) {
			for _, id := range fcIds {
				t.Run("FC ID "+id, func(t *testing.T) {
					fc, err := s.FetchFreeCompany(id)
					if err != nil {
						if lang == SiteLang("zh") {
							return
						}

						t.Errorf(err.Error())
					}

					failIfStringEmpty(t, "FC active state", string(fc.Active))
					failIfStringEmpty(t, "FC name", fc.Name)
					failIfStringEmpty(t, "FC ID", fc.ID)
					failIfStringEmpty(t, "FC world", fc.DC)
					failIfStringEmpty(t, "FC DC", fc.DC)
					failIfStringEmpty(t, "FC recruitment status", string(fc.Recruitment))
					failIfOlderThanGameRelease(t, "FC formation", fc.Formed)
				})
			}
		})
	}
}

func TestFetchFreeCompanyMembers(t *testing.T) {
	for _, lang := range langCodes {
		s := NewScraper(lang)

		t.Run("SiteLang: "+string(lang), func(t *testing.T) {
			for _, id := range fcIds {
				t.Run("FC ID "+id, func(t *testing.T) {
					for member := range s.FetchFreeCompanyMembers(id) {
						if member.Error != nil {
							if lang == SiteLang("zh") {
								return
							}

							t.Errorf(member.Error.Error())
						}

						failIfStringEmpty(t, "Member avatar", member.Avatar)
						failIfNumberZero(t, "Member ID", int64(member.ID))
						failIfStringEmpty(t, "Member name", member.Name)
						failIfStringEmpty(t, "Member world", member.World)
						failIfStringEmpty(t, "Member DC", member.DC)
					}
				})
			}
		})
	}
}

func TestSearchFreeCompanies(t *testing.T) {
	for _, lang := range langCodes {
		s := NewScraper(lang)

		t.Run("SiteLang: "+string(lang), func(t *testing.T) {
			opts := search.FreeCompanyOptions{}

			for fc := range s.SearchFreeCompanies(opts) {
				if fc.Error != nil {
					if lang == SiteLang("zh") {
						return
					}

					t.Errorf(fc.Error.Error())
				}

				failIfStringEmpty(t, "FC active state", string(fc.Active))
				failIfNumberZero(t, "FC active members", int64(fc.ActiveMembers))
				failIfGCInvalid(t, "FC Grand Company", fc.GrandCompany)
				failIfStringEmpty(t, "FC ID", fc.ID)
				failIfStringEmpty(t, "FC name", fc.Name)
				failIfStringEmpty(t, "FC world", fc.World)
				failIfStringEmpty(t, "FC DC", fc.DC)
				failIfStringEmpty(t, "FC estate", fc.Estate)
				failIfOlderThanGameRelease(t, "FC formed", fc.Formed)
			}
		})
	}
}

func TestSearchCharacters(t *testing.T) {
	for _, lang := range langCodes {
		s := NewScraper(lang)

		t.Run("SiteLang: "+string(lang), func(t *testing.T) {
			opts := search.CharacterOptions{}

			for character := range s.SearchCharacters(opts) {
				if character.Error != nil {
					if lang == SiteLang("zh") {
						return
					}

					t.Errorf(character.Error.Error())
				}

				failIfStringEmpty(t, "Character avatar", character.Avatar)
				failIfNumberZero(t, "Character ID", int64(character.ID))
				failIfStringEmpty(t, "Character name", character.Name)
				failIfStringEmpty(t, "Character world", character.World)
				failIfStringEmpty(t, "Character DC", character.DC)
			}
		})
	}
}

func TestSearchCWLS(t *testing.T) {
	for _, lang := range langCodes {
		s := NewScraper(lang)

		t.Run("SiteLang: "+string(lang), func(t *testing.T) {
			opts := search.CWLSOptions{}

			for cwls := range s.SearchCWLS(opts) {
				if cwls.Error != nil {
					if lang == SiteLang("zh") {
						return
					}

					t.Errorf(cwls.Error.Error())
				}

				failIfStringEmpty(t, "CWLS ID", cwls.ID)
				failIfStringEmpty(t, "CWLS name", cwls.Name)
				failIfStringEmpty(t, "CWLS DC", cwls.DC)
				failIfNumberZero(t, "CWLS active members", int64(cwls.ActiveMembers))
			}
		})
	}
}

func TestSearchLinkshells(t *testing.T) {
	for _, lang := range langCodes {
		s := NewScraper(lang)

		t.Run("SiteLang: "+string(lang), func(t *testing.T) {
			opts := search.LinkshellOptions{}

			for ls := range s.SearchLinkshells(opts) {
				if ls.Error != nil {
					if lang == SiteLang("zh") {
						return
					}

					t.Errorf(ls.Error.Error())
				}

				failIfStringEmpty(t, "Linkshell ID", ls.ID)
				failIfStringEmpty(t, "Linkshell name", ls.Name)
				failIfStringEmpty(t, "Linkshell world", ls.World)
				failIfStringEmpty(t, "Linkshell DC", ls.DC)
				failIfNumberZero(t, "Linkshell active members", int64(ls.ActiveMembers))
			}
		})
	}
}

func TestSearchPVPTeams(t *testing.T) {
	for _, lang := range langCodes {
		s := NewScraper(lang)

		t.Run("SiteLang: "+string(lang), func(t *testing.T) {
			opts := search.PVPTeamOptions{}

			for ls := range s.SearchPVPTeams(opts) {
				if ls.Error != nil {
					if lang == SiteLang("zh") {
						return
					}

					t.Errorf(ls.Error.Error())
				}

				failIfStringEmpty(t, "PVP team ID", ls.ID)
				failIfStringEmpty(t, "PVP team name", ls.Name)
				failIfStringEmpty(t, "PVP team DC", ls.DC)
			}
		})
	}
}
