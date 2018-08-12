package database

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/postgres"
)

type Kata struct {
	gorm.Model
	Uid string `gorm:"primary_key;unique"`
	Title string
	Kyu int
	Url string
	Data postgres.Jsonb
  }

type Client struct {
	Db *gorm.DB
}

func NewClient(username string, password string, dbName string, host string, port int) *Client {
	dbInfo := fmt.Sprintf(
		"user=%s password=%s host=%s port=%d dbname=%s sslmode=disable",
		username,
		password,
		host,
		port,
		dbName,
	)

	log.Printf("Connecting to database `%s` on host `%s:%d` with user `%s`.", dbName, host, port, username)
	dbClient, err := gorm.Open("postgres", dbInfo)
	if err != nil {
		panic(err)
	}

	log.Println("Migrating DB schema.")
	dbClient.AutoMigrate(&Kata{})

	return &Client{
		Db: dbClient,
	}
}

func (c *Client) InsertKatas(kataChan chan map[string]interface{}) {
	for {
		var kata Kata
		kataMap := <- kataChan

		uid, _ := kataMap["uid"].(string)
		title, _ := kataMap["title"].(string)
		kyu, _ := kataMap["kyu"].(int)
		url, _ := kataMap["url"].(string)

		log.Printf("Inserting Kata to DB with vals: %v", kataMap)

		c.Db.Where(Kata{
			Uid: uid,
		}).Assign(Kata{
			Title: title,
			Kyu: kyu,
			Url: url,
		}).FirstOrCreate(&kata)
	}
}
