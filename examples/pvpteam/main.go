package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/karashiiro/godestone"
)

func main() {
	s, err := godestone.NewScraper()
	if err != nil {
		log.Fatalln(err)
	}

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
