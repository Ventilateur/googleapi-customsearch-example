package main

import (
	"fmt"
	"googlesearch"
)

const (
	SEARCH_ENGINE = "Your search engine"	// To be configured personally
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}


func main() {
	searcher, err := googlesearch.New(API_KEY, SEARCH_ENGINE, true)
	check(err)
	
	results, err := searcher.Search("apples", 10)
	check(err)
	
	var imgLinks []string
	for _, item := range(results.Items) {
		imgLinks = append(imgLinks, item.Link)
	}

	for _, link := range(imgLinks) {
		fmt.Println(link)
	}
}
