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
