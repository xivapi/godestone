package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/karashiiro/godestone"
)

func main() {
	s := godestone.NewScraper(godestone.DE)

	id, err := strconv.ParseUint(os.Args[1], 10, 32)
	if err != nil {
		log.Fatalln(err)
	}

	m, err := s.FetchCharacterMinions(uint32(id))
	if err != nil {
		log.Fatalln(err)
	}

	for _, minion := range m {
		log.Println(minion.Name, fmt.Sprintf("(ID: %d)", minion.ID))
	}
}
