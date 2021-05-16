package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

func ParseHtmlFromUrl(url string) string {
	fmt.Printf("[*] Parsing HTML from %s ...\n", url)

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

// Reads a two-dimensional slice of strings and turn it into a xlsx file.
// Returns the url where the file can be downloaded.
func LinksToXlsx(filename string, links [][]string) string {
	// Create new xlsx file
	return "Hello world"
}

// Creates an Excel file on a given directory
func CreateExcel(f *excelize.File, dir string) *excelize.File {
	if err := f.SaveAs(dir); err != nil {
		fmt.Println(err)
	}

	return f
}

// Creates a row with data, on a given sheet
func CreateSheet(f *excelize.File, sheetName string, data [][]map[string]string) *excelize.File {
	if f == nil {
		f = excelize.NewFile()
	}

	index := f.NewSheet(sheetName)
	f.SetActiveSheet(index)

	for _, col := range data {
		for _, row := range col {
			for index, value := range row {
				f.SetCellValue(sheetName, index, value)
			}
		}
	}

	return f
}

// Receives a string map and multiple string values as input, and appends data to
// the last available rows.
func ArrayForExcel(file [][]map[string]string, row ...string) [][]map[string]string {
	var data []map[string]string
	var ExcelRow map[string]string = make(map[string]string)

	col := []string{
		"A", "B", "C", "D", "E", "F", "G",
		"H", "I", "J", "K", "L", "M", "N",
		"O", "P", "Q", "R", "S", "T", "U",
		"X", "Y", "Z"}

	for index, value := range row {
		ExcelCol := col[index] + fmt.Sprint(len(file)+1)
		ExcelRow[ExcelCol] = value
		data = append(data, ExcelRow)
	}

	file = append(file, data)

	return file
}

func HandleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
