/*
1. Parse this CSV file, putting two fields from
the contents of the CSV file into a data structure.
2. Parse an html template, then pass the data from step 1 into the CSV template; have the html template display the CSV data in a web page.
*/
package main

import (
	"encoding/csv"
	"html/template"
	"log"
	"os"
	"strconv"
	"time"
)

var tpl *template.Template

/*
Record ...
Date - time.Time
Open - float64
*/
type Record struct {
	Date time.Time
	Open float64
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func prs(filePath string) []Record {
	//open file
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	/*
		r.Read() 	- will only get the first row
		r.ReadAll() - will get every row
	*/
	rows, err := r.ReadAll()
	if err != nil {
		log.Fatalln(err)
	}

	// make an empty array with the size of rows
	records := make([]Record, 0, len(rows))

	//looping through each row
	for i, row := range rows {
		// ignoring the csv files head
		if i == 0 {
			continue
		}

		// put each first two rows into own slices
		date, _ := time.Parse("2006-01-02", row[0])
		open, _ := strconv.ParseFloat(row[1], 64)

		//append slices to records and finally returns it
		records = append(records, Record{
			Date: date,
			Open: open,
		})
	}
	return records
}

func main() {

	// parse csv
	records := prs("data.csv")

	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", records)
	if err != nil {
		log.Fatalln(err)
	}

}
