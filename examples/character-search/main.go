package main

import (
	"log"
	"os"

	"github.com/xivapi/godestone"
)

func main() {
	s := godestone.NewScraper(godestone.EN)

	opts := godestone.CharacterOptions{}
	if len(os.Args) > 1 {
		opts = godestone.CharacterOptions{
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
