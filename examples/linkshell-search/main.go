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

	opts := godestone.SearchLinkshellOptions{
		Name: os.Args[1],
	}

	for ls := range s.SearchLinkshells(opts) {
		if ls.Error != nil {
			log.Fatalln(ls.Error)
		}

		log.Println(ls)
	}
}
