package main

import (
	"log"
	"os"

	"github.com/karashiiro/godestone"
)

func main() {
	s, err := godestone.NewScraper()
	if err != nil {
		log.Fatalln(err)
	}

	opts := godestone.SearchCharacterOptions{
		Name:  os.Args[1] + " " + os.Args[2],
		World: os.Args[3],
	}

	for character := range s.SearchCharacters(opts) {
		if character.Error != nil {
			log.Fatalln(character.Error)
		}

		log.Println(character)
	}
}
