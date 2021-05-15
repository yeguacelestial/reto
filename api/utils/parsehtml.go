package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func ParseHtmlFromUrl(url string) string {
	fmt.Printf("Parsing HTML from %s ...\n", url)

	res, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	// Reads HTML as slice of bytes
	html, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	return string(html)
}
