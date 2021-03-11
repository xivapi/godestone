package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/xivapi/godestone"
)

func main() {
	s := godestone.NewScraper(godestone.EN)

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
