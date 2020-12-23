package main

import (
	"log"
	"os"

	"github.com/karashiiro/godestone"
)

func main() {
	s, err := godestone.NewScraper()
	if err != nil {
		log.Fatalln(err)
	}

	opts := godestone.SearchPVPTeamOptions{
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
