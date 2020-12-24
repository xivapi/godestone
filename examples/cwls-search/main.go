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
