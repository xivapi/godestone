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

	opts := godestone.SearchCWLSOptions{
		Name: os.Args[1],
	}

	for cwls := range s.SearchCWLS(opts) {
		if cwls.Error != nil {
			log.Fatalln(cwls.Error)
		}

		log.Println(cwls)
	}
}
