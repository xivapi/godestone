package main

import (
	"log"
	"os"

	"github.com/karashiiro/godestone"
	"github.com/karashiiro/godestone/search"
)

func main() {
	s := godestone.NewScraper(godestone.EN)

	opts := search.CWLSOptions{
		Name: os.Args[1],
	}

	for cwls := range s.SearchCWLS(opts) {
		if cwls.Error != nil {
			log.Fatalln(cwls.Error)
		}

		log.Println(cwls)
	}
}
