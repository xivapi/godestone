package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/xivapi/godestone"
)

func main() {
	s := godestone.NewScraper(godestone.EN)

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
