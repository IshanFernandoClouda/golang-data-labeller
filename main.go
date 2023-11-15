package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

var CATEGORY int = 0
var PRODUCT_ID int = 0
var numberOfCols int = 0

func main() {
	filePath := "bisco4mill.csv"

	// Open the CSV file
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	// Create a CSV reader
	reader := csv.NewReader(file)

	// Read all the records from CSV
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for index, col := range records[0] {
		if col == "_CategoryName" {
			CATEGORY = index
		} else if col == "_SkuId (Not changeable)" {
			PRODUCT_ID = index
		}
	}
	numberOfCols = len(records[0])
	records = records[1:]

	processStrings(0, 0, len(records), records)

}

func processStrings(id int, startIndex int, endIndex int, records [][]string) {
	// for each in records
	for _, record1 := range records[startIndex:endIndex] {
		category1 := record1[CATEGORY]
		product_id1 := record1[PRODUCT_ID]

		if len(record1) != numberOfCols {
			fmt.Println("Wrong line")
		}
		for _, record2 := range records {
			category2 := record2[CATEGORY]
			product_id2 := record2[PRODUCT_ID]

			if category1 == category2 {
				fmt.Println(product_id1 + "//" + product_id2 + "//" + "1")
			} else {
				fmt.Println(product_id1 + "//" + product_id2 + "//" + "0")
			}
		}
	}
}
