package main

import (
	"log"
	"os"

	"github.com/xivapi/godestone/v2"
)

func main() {
	s := godestone.NewScraper(godestone.EN)

	for member := range s.FetchFreeCompanyMembers(os.Args[1]) {
		if member.Error != nil {
			log.Fatalln(member.Error)
		}

		log.Println(member)
	}
}
