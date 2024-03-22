// Started with example for reading csv files from the following website, and
// made significant changes to suit my use case:
// https://gosamples.dev/read-csv/
package utils

import (
	"encoding/csv"
	"log"
	"os"
)

type Text struct {
	key string
	val string
}

func CsvToTextList(file_path string) map[string]string {
	// open file
	f, err := os.Open(file_path)
	if err != nil {
		log.Fatal(err)
	}

	// remember to close the file at the end of the program
	defer f.Close()

	// read csv values using csv.Reader
	csvReader := csv.NewReader(f)
	// use a custom character for separation of key/value in line
	csvReader.Comma = '|'
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// initialize the map, then read csv into map
	text_map := make(map[string]string)
	for i, line := range data {
		if i > 0 { // omit header line
			text_map[line[0]] = line[1]
		}
	}
	return text_map
	// return text_list
}

// func csvToData(file_name string) [][]string {
// 	// open file
// 	f, err := os.Open("data.csv")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
//
// 	// remember to close the file at the end of the program
// 	defer f.Close()
//
// 	// read csv values using csv.Reader
// 	csvReader := csv.NewReader(f)
// 	data, err := csvReader.ReadAll()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return data
// }
