package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/karashiiro/godestone"
)

func main() {
	s := godestone.NewScraper(godestone.EN)

	c, err := s.FetchCWLS(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}

	cJSON, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(cJSON))
}
