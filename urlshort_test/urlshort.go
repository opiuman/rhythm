package main

import (
	"flag"
	"fmt"
	"github.com/opiuman/rhythm"
	"log"
)

func main() {
	longUrl := flag.String("longUrl", "http://www.google.com", "url")
	token := flag.String("token", "b5e891db167ab4b65974cb3927348f35e2889094", "token")
	flag.Parse()
	shortUrl, err := rhythm.ShortUrl(*longUrl, *token)
	if err != nil {
		log.Printf("shortener Url failed: %s", err)
	}
	fmt.Printf("shortUrl= %s \n", shortUrl)
}
