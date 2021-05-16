package utils

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
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

// Reads a two-dimensional slice of strings and turn it into a csv file.
// Returns the url where the file can be downloaded.
func LinksToCsv(filename string, links [][]string) string {
	file, err := os.Create(filename)
	HandleErr(err)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, value := range links {
		err := writer.Write(value)
		HandleErr(err)
	}

	return filename
}

func HandleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
