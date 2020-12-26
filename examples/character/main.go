package main

import (
	"encoding/json"
	"log"
	"os"
	"strconv"

	"github.com/karashiiro/godestone"
)

func main() {
	s := godestone.NewScraper(godestone.EN)

	id, err := strconv.ParseUint(os.Args[1], 10, 32)
	if err != nil {
		log.Fatalln(err)
	}

	c, err := s.FetchCharacter(uint32(id))
	if err != nil {
		log.Fatalln(err)
	}

	cJSON, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(cJSON))
}
