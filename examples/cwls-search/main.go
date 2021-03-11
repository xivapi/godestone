package main

import (
	"log"
	"os"

	"github.com/xivapi/godestone"
)

func main() {
	s := godestone.NewScraper(godestone.EN)

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
