package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/xivapi/godestone"
)

func main() {
	s := godestone.NewScraper(godestone.EN)

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
