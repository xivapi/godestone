package main

import (
	"log"
	"os"

	"github.com/karashiiro/godestone"
	"github.com/karashiiro/godestone/search"
)

func main() {
	s, err := godestone.NewScraper(godestone.EN)
	if err != nil {
		log.Fatalln(err)
	}

	opts := search.PVPTeamOptions{
		Name: os.Args[1],
	}

	for pvpTeam := range s.SearchPVPTeams(opts) {
		if pvpTeam.Error != nil {
			log.Fatalln(pvpTeam.Error)
		}

		log.Println(pvpTeam)
		log.Println(pvpTeam.CrestLayers)
	}
}
