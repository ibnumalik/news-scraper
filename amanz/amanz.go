package amanz

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/gocolly/colly/v2"
)

const (
	defaultOutput string = "news.json"
	url           string = "https://amanz.my"
)

var newstype string = "latest"

func usage() {
	fmt.Println(`Usage: ./amanzscraper [newstype]

[newstype]	can be any one of these: trending, latest, featured

.`)
	os.Exit(1)
}

func Get() {
	if len(os.Args) > 2 {
		usage()
	}

	if len(os.Args) > 1 {
		newstype = os.Args[1]
	}

	var output string = defaultOutput

	var news []string

	switch newstype {
	case "trending":
		news = Trending()
	case "latest":
		news = Latest()
	case "featured":
		news = Featured()
	default:
		fmt.Println("Cannot proceed.")
		os.Exit(1)
	}

	// Read the news in the stdout. Maybe we can add option to read or to save.
	for i, new := range news {
		fmt.Println(i+1, new)
	}

	toWrite, err := json.MarshalIndent(news, "", "")
	if err != nil {
		fmt.Printf("Cant't marshall: %v", err.Error())
		os.Exit(1)
	}

	err = ioutil.WriteFile(output, toWrite, 0644)
	if err != nil {
		fmt.Printf("Can't write to file: %v", err.Error())
		os.Exit(1)
	}
}

func Featured() []string {
	var news []string

	c := colly.NewCollector(
		colly.AllowedDomains("amanz.my"),
	)

	c.OnHTML(".home-features", func(e *colly.HTMLElement) {
		titles := e.ChildTexts(".f5 .featured-meta h4")
		for _, title := range titles {
			news = append(news, title)
		}
	})

	c.Visit(url)

	return news
}

func Latest() []string {
	var news []string

	c := colly.NewCollector(
		colly.AllowedDomains("amanz.my"),
	)

	c.OnHTML("div.no-mobile > .row > .twelve[x-show*=terkini]", func(e *colly.HTMLElement) {
		titles := e.ChildTexts(".article h5")
		for _, title := range titles {
			news = append(news, title)
		}
	})

	c.Visit(url)

	return news
}

func Trending() []string {
	var news []string

	c := colly.NewCollector(
		colly.AllowedDomains("amanz.my"),
	)

	c.OnHTML("div.no-mobile > .row > .twelve[x-show*=sohor]", func(e *colly.HTMLElement) {
		titles := e.ChildTexts(".article h5")
		for _, title := range titles {
			news = append(news, title)
		}
	})

	c.Visit(url)

	return news
}
