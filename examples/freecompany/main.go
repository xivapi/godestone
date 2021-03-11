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

	fc, err := s.FetchFreeCompany(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}

	fcJSON, err := json.MarshalIndent(fc, "", "  ")
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(fcJSON))
}
