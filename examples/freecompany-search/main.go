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

	opts := godestone.SearchFreeCompanyOptions{
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
