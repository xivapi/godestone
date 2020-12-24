package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/karashiiro/godestone"
)

func main() {
	s, err := godestone.NewScraper(godestone.EN)
	if err != nil {
		log.Fatalln(err)
	}

	l, err := s.FetchLinkshell(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}

	lJSON, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(lJSON))
}
