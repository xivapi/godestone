package main

import (
	"log"
	"os"

	"github.com/karashiiro/bingode"
	"github.com/xivapi/godestone/v2"
)

func main() {
	s := godestone.NewScraper(bingode.New(), godestone.EN)

	opts := godestone.CWLSOptions{
		Name: os.Args[1],
	}

	for cwls := range s.SearchCWLS(opts) {
		if cwls.Error != nil {
			log.Fatalln(cwls.Error)
		}

		log.Println(cwls)
	}
}
