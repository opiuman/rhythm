package main

import (
	"flag"
	"fmt"
	"github.com/opiuman/rhythm"
)

func main() {
	longUrl := flag.String("longUrl", "http://www.google.com", "url")
	flag.Parse()
	//fmt.Printf(*longUrl)
	shortUrl, err := rhythm.ShortUrl(*longUrl)
	if err != nil {
		fmt.Errorf("shortener Url failed: %s", err)
	}
	fmt.Printf("shortUrl= %s \n", shortUrl)
}
