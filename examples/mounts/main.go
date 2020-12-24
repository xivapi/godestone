package main

import (
	"log"
	"os"
	"strconv"

	"github.com/karashiiro/godestone"
)

func main() {
	s, err := godestone.NewScraper(godestone.EN)
	if err != nil {
		log.Fatalln(err)
	}

	id, err := strconv.ParseUint(os.Args[1], 10, 32)
	if err != nil {
		log.Fatalln(err)
	}

	m, err := s.FetchCharacterMounts(uint32(id))
	if err != nil {
		log.Fatalln(err)
	}

	for _, mount := range m {
		log.Println(mount.Name)
	}
}
