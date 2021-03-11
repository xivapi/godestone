package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/karashiiro/bingode"
	"github.com/xivapi/godestone/v2"
)

func main() {
	s := godestone.NewScraper(bingode.New(), godestone.EN)

	p, err := s.FetchPVPTeam(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}

	pJSON, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(pJSON))
}
