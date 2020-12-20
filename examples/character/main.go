package main

import (
	"log"
	"os"
	"strconv"
	"text/template"

	"github.com/karashiiro/godestone"
)

func main() {
	s, err := godestone.NewScraper()
	if err != nil {
		log.Fatalln(err)
	}

	id, err := strconv.ParseUint(os.Args[1], 10, 32)
	if err != nil {
		log.Fatalln(err)
	}

	c, err := s.FetchCharacter(uint32(id))
	if err != nil {
		log.Fatalln(err)
	}

	tmpl, err := template.New("character").Parse(
		`{ 
	Avatar:          "{{.Avatar}}",
	Bio:             "{{.Bio}}",
	DC:              "{{.DC}}",
	FreeCompanyID:   "{{.FreeCompanyID}}",
	FreeCompanyName: "{{.FreeCompanyName}}",
	Gender:          {{.Gender}},
	GuardianDeity:   {{.GuardianDeity}},
	Name:            "{{.Name}}",
	Nameday:         "{{.Nameday}}",
	Portrait:        "{{.Portrait}}",
	PvPTeamID:       "{{.PvPTeamID}}",
	Race:            {{.Race}},
	Server:          "{{.Server}}",
	Town:            {{.Town}},
	Tribe:           {{.Tribe}}
}`)
	if err != nil {
		log.Fatalln(err)
	}

	tmpl.Execute(os.Stdout, c)
}
