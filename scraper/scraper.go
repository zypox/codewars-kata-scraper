package scraper

import (
	"log"
	"net/http"
	"sync"

	"github.com/PuerkitoBio/goquery"

	"github.com/zygimantasp/codewars-kata-scraper/helper"
)

func getHTML(url *string) *goquery.Document {
	res, err := http.Get(*url)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("Unsuccessful HTTP response status code: %d %s", res.StatusCode, res.Status)
	}
	
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	return doc
}

func GetKataCount(url string) int {
	html := getHTML(&url)
	kataCountStr := html.Find(".items-list > .prn > p.mtn").Text()

	return helper.ParseFirstDigit(&kataCountStr)
}

func parseKyu(kataBlock *goquery.Selection) int {
	kyuStr := kataBlock.Find(".item-title > .mrm span").Text()
	return helper.ParseFirstDigit(&kyuStr)
}

func parseLanguages(kataBlock *goquery.Selection) []string {
	var languages []string
	kataBlock.Find(".language-icons > li > a").Each(func(i int, langBlock *goquery.Selection) {
		lang, _ := langBlock.Attr("data-language")
		languages = append(languages, lang)
	})
	return languages
}

func parseKeywordTags(kataBlock *goquery.Selection) []string {
	var keywordTags []string
	kataBlock.Find(".keyword-tag > a").Each(func(i int, keywordTagEl *goquery.Selection) {
		keywordTag := keywordTagEl.Text()
		keywordTags = append(keywordTags, keywordTag)
	})
	return keywordTags
}

func ScrapeKataPage(url *string, wg *sync.WaitGroup, katasChan chan map[string]interface{}) {
	defer wg.Done()
	html := getHTML(url)
	html.Find(".items-list .kata").Each(func(i int, kataBlock *goquery.Selection) {
		uid, _ := kataBlock.Attr("id")
		title := kataBlock.Find(".item-title a").Text()
		url, _ := kataBlock.Find(".item-title a").Attr("href")
		kyu := parseKyu(kataBlock)
		languages := parseLanguages(kataBlock)
		keywordTags := parseKeywordTags(kataBlock)
		kataParams := map[string]interface{}{
			"uid": uid,
			"title": title,
			"url": url,
			"kyu": kyu,
			"languages": languages,
			"keywordTags": keywordTags,
		}
		katasChan <- kataParams
	})
}
