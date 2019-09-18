package main

import (
	"log"
	"rss-twitter/config"

	"github.com/joho/godotenv"
)

var conf = config.New("config.json")

func init() {
	var err error

	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	err = conf.Load()
	if err != nil {
		log.Fatal("Error loading config file")
	}
}

func main() {
	for _, url := range conf.RssFeeds {
		err := processRSS(url)
		if err != nil {
			log.Printf("Error processing %s - %+v", url, err)
		}
	}

	err := conf.Save()
	if err != nil {
		log.Printf("Error saving config: %+v", err)
	}
}
