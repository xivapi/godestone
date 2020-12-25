package main

import (
	"log"
	"os"

	"github.com/karashiiro/godestone"
	"github.com/karashiiro/godestone/search"
)

func main() {
	s, err := godestone.NewScraper(godestone.EN)
	if err != nil {
		log.Fatalln(err)
	}

	opts := search.CharacterOptions{}
	if len(os.Args) > 1 {
		opts = search.CharacterOptions{
			Name:  os.Args[1] + " " + os.Args[2],
			World: os.Args[3],
		}
	}

	for character := range s.SearchCharacters(opts) {
		if character.Error != nil {
			log.Fatalln(character.Error)
		}

		log.Println(character)
	}
}
