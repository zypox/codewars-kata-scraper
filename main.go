package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"sync"

	"github.com/zygimantasp/codewars-kata-scraper/helper"
	"github.com/zygimantasp/codewars-kata-scraper/database"
	"github.com/zygimantasp/codewars-kata-scraper/scraper"
)

const KATA_COUNT_PER_PAGE = 30

func main() {
	log.Println("-----STARTING CODEWARS KATA SCRAPER-----")
	configPath := flag.String("config", "conf", "path to config dir")
	dbConfig := helper.ReadFromConfigOrPanic(*configPath, "database")

	dbClient := database.NewClient(
		dbConfig.GetString("postgresql.connection.username"),
		dbConfig.GetString("postgresql.connection.password"),
		dbConfig.GetString("postgresql.connection.database"),
		dbConfig.GetString("postgresql.connection.host"),
		dbConfig.GetInt("postgresql.connection.port"),
	)

	kataCount := scraper.GetKataCount("https://www.codewars.com/kata/latest?page=1")
	pageCount := int(math.Ceil(float64(kataCount) / float64(KATA_COUNT_PER_PAGE)))
	log.Printf("Total kata and page count: `%d` and `%d`.", kataCount, pageCount)

	wg := &sync.WaitGroup{}
	katasChan := make(chan map[string]interface{})

	for i := 0; i < pageCount; i++ {
		url := fmt.Sprintf("https://www.codewars.com/kata/latest?page=%d", i)
		wg.Add(1)
		go scraper.ScrapeKataPage(&url, wg, katasChan)
	}

	go dbClient.InsertKatas(katasChan)

	wg.Wait()
}
