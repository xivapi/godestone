package main

import (
	"log"
	"os"

	"github.com/karashiiro/godestone"
)

func main() {
	s := godestone.NewScraper(godestone.EN)

	opts := godestone.LinkshellOptions{
		Name: os.Args[1],
	}

	for ls := range s.SearchLinkshells(opts) {
		if ls.Error != nil {
			log.Fatalln(ls.Error)
		}

		log.Println(ls)
	}
}
