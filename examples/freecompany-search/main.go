package main

import (
	"encoding/json"
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

	opts := search.FreeCompanyOptions{
		Name: os.Args[1],
	}

	for fc := range s.SearchFreeCompanies(opts) {
		if fc.Error != nil {
			log.Fatalln(fc.Error)
		}

		fcJSON, err := json.MarshalIndent(fc, "", "  ")
		if err != nil {
			log.Fatalln(err)
		}

		log.Println(string(fcJSON))
	}
}
