package main

import (
	"log"

	"github.com/karashiiro/godestone"
)

func main() {
	s, err := godestone.NewScraper()
	if err != nil {
		log.Fatalln(err)
	}

	c, err := s.FetchCharacter(2831882)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(c.Avatar)
}
