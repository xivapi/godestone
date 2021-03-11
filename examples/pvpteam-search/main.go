package main

import (
	"log"
	"os"

	"github.com/xivapi/godestone"
)

func main() {
	s := godestone.NewScraper(godestone.EN)

	opts := godestone.PVPTeamOptions{
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
