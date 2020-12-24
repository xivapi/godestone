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

	for achievement := range s.FetchCharacterAchievements(uint32(id)) {
		if achievement.Error != nil {
			log.Fatalln(achievement.Error)
		}
		log.Println(achievement)
		log.Println(achievement.TotalAchievementInfo)
	}
}
