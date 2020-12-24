package main

import (
	"log"
	"os"

	"github.com/karashiiro/godestone"
)

func main() {
	s, err := godestone.NewScraper(godestone.EN)
	if err != nil {
		log.Fatalln(err)
	}

	for member := range s.FetchFreeCompanyMembers(os.Args[1]) {
		if member.Error != nil {
			log.Fatalln(member.Error)
		}

		log.Println(member)
	}
}
