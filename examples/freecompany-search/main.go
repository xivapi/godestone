package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/xivapi/godestone"
)

func main() {
	s := godestone.NewScraper(godestone.EN)

	opts := godestone.FreeCompanyOptions{}
	if len(os.Args) > 1 {
		opts = godestone.FreeCompanyOptions{
			Name: os.Args[1],
		}
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
