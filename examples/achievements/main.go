package main

import (
	"log"
	"os"
	"strconv"

	"github.com/xivapi/godestone/v2"
)

func main() {
	s := godestone.NewScraper(godestone.EN)

	id, err := strconv.ParseUint(os.Args[1], 10, 32)
	if err != nil {
		log.Fatalln(err)
	}

	a, aai, err := s.FetchCharacterAchievements(uint32(id))
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(aai)
	for _, achievement := range a {
		log.Println(achievement)
	}
}
