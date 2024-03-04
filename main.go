package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.Open("data.csv")
	if err != nil {
		log.Fatal(err)
	}

	// remember to close the file at the end of the program
	defer f.Close()

	// read csv values using csv.Reader
	csvReader := csv.NewReader(f)
	query := `INSERT INTO data_scope
	(scope, description)
VALUES `
	for {
		scope := "customer_category_id:"
		description := "Data Scopes customer_category_id With ID "
		rec, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		// do something with read line
		query += "('" + scope + rec[0] + "','" + description + rec[0] + "'),"
		fmt.Printf("%+v\n", rec)
	}

	log.Fatal("DONE")
}
