package main

import (
	"encoding/csv"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	inputFile, err := os.Open("input.csv")
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer inputFile.Close()

	reader := csv.NewReader(inputFile)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}

	// Find max values for normalization
	maxDl, maxNos, maxNod, maxNoh := 0, 0, 0, 0
	for _, record := range records[1:] {
		dl, _ := strconv.Atoi(record[2])
		nos, _ := strconv.Atoi(record[3])
		nod, _ := strconv.Atoi(record[4])
		noh, _ := strconv.Atoi(record[5])

		maxDl = int(math.Max(float64(maxDl), float64(dl)))
		maxNos = int(math.Max(float64(maxNos), float64(nos)))
		maxNod = int(math.Max(float64(maxNod), float64(nod)))
		maxNoh = int(math.Max(float64(maxNoh), float64(noh)))
	}

	// Normalize numeric columns and one-hot encode TLD column
	tldSet := make(map[string]bool)
	newRecords := [][]string{}
	newRecords = append(newRecords, []string{"td", "tld", "dl", "nos", "nod", "noh", "m"})
	for _, record := range records[1:] {
		tld := record[1]
		tldSet[tld] = true

		dl, _ := strconv.Atoi(record[2])
		nos, _ := strconv.Atoi(record[3])
		nod, _ := strconv.Atoi(record[4])
		noh, _ := strconv.Atoi(record[5])

		newRecords = append(newRecords, []string{
			record[0],
			tld,
			fmt.Sprintf("%.2f", float64(dl)/float64(maxDl)),
			fmt.Sprintf("%.2f", float64(nos)/float64(maxNos)),
			fmt.Sprintf("%.2f", float64(nod)/float64(maxNod)),
			fmt.Sprintf("%.2f", float64(noh)/float64(maxNoh)),
			record[6],
		})
	}

	// One-hot encoding of TLD column
	tlds := make([]string, 0, len(tldSet))
	for tld := range tldSet {
		tlds = append(tlds, tld)
	}

	encodedHeader := []string{"td"}
	encodedHeader = append(encodedHeader, tlds...)
	encodedHeader = append(encodedHeader, "dl", "nos", "nod", "noh", "m")

	encodedRecords := [][]string{encodedHeader}
	for _, record := range newRecords[1:] {
		encodedRecord := []string{record[0]}
		for _, tld := range tlds {
			if record[1] == tld {
				encodedRecord = append(encodedRecord, "1")
			} else {
				encodedRecord = append(encodedRecord, "0")
			}
		}
		encodedRecord = append(encodedRecord, record[2:]...)
		encodedRecords = append(encodedRecords, encodedRecord)
	}

	// Write the preprocessed data to a new CSV file
	outputFile, err := os.Create("preprocessed_data.csv")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()

	writer := csv.NewWriter(outputFile)
	err = writer.WriteAll(encodedRecords)
	if err != nil {
		fmt.Println("Error writing to output file:", err)
		return
	}

	fmt.Println("Preprocessing completed successfully!")
}
