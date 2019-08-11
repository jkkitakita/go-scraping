package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"reflect"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/go-yaml/yaml"
	"github.com/gocarina/gocsv"
	"github.com/jkkitakita/go-scraping/domain"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

func failOnError(err error) {
	if err != nil {
		log.Fatal("Error:", err)
	}
}

// Scrape returns entity
func Scrape(site domain.Site, url string) domain.SakeEntity {
	fmt.Printf("url[%%+v] -> %+v\n", url)
	// Request the HTML page.
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var entity domain.SakeEntity

	entity.URL = url
	entity.Category = site.Category.Name

	doc.Find(".spec .name").Each(func(i int, s *goquery.Selection) {
		entity.Name = s.Text()
	})

	doc.Find(".nomal_price .price").Each(func(i int, s *goquery.Selection) {
		entity.Price = s.Text()
	})

	getValues(doc, &entity)

	return entity
}

func getValues(doc *goquery.Document, entity *domain.SakeEntity) {
	doc.Find(".deco_table tbody tr").Each(func(i int, s *goquery.Selection) {
		t := reflect.TypeOf(domain.SakeEntity{})
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			j := field.Tag.Get("csv")
			if j == s.Find("th").Text() {
				reflect.ValueOf(entity).Elem().FieldByName(field.Name).SetString(strings.TrimSpace(s.Find("td").Text()))
			}
		}
	})
}

// Pagelist returns list of url
func Pagelist(site domain.Site) []string {

	urllists := []string{}
	fmt.Printf("site[%%+v] -> %+v\n", site)
	for i := site.Pagination.Page.Offset; i <= site.Pagination.Page.Limit; i++ {
		//カテゴリ一覧から商品URLのリストを返す
		res, err := http.Get(fmt.Sprintf(
			`%+v%+v%+v%+v%+v%+v`,
			site.Domain,
			site.Category.Path,
			site.Pagination.PerPage.Suffix,
			site.Pagination.PerPage.Number,
			site.Pagination.Page.Suffix,
			i,
		))
		if err != nil {
			log.Fatal(err)
		}

		defer res.Body.Close()
		if res.StatusCode != 200 {
			log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
		}

		//楽天はEUC-JPなのでエンコードする
		utfBody := transform.NewReader(bufio.NewReader(res.Body), japanese.EUCJP.NewDecoder())
		//ページ読み込み

		doc, err := goquery.NewDocumentFromReader(utfBody)

		//型番
		doc.Find(".itemlist .name a").Each(func(i int, s *goquery.Selection) {
			var url, _ = s.Attr("href")
			urllists = append(urllists, url)
		})
	}
	fmt.Printf("the number of urllists -> %+v\n", len(urllists))

	return urllists
}

func main() {
	var site domain.Site

	// test.yamlを []byte として読み込みます。
	buf, err := ioutil.ReadFile("./category.yaml")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = yaml.Unmarshal(buf, &site)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	urls := Pagelist(site)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	entities := []domain.SakeEntity{}
	for _, url := range urls {
		entities = append(entities, Scrape(site, url))
	}

	filename := fmt.Sprintf(`%+v_%+v.csv`, site.Category.Name, time.Now().Format("2006-01-02"))
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0600)
	defer file.Close()
	gocsv.MarshalFile(&entities, file)
}
