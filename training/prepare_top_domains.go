package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	csvFile, err := os.Open("majestic_million.csv")
	if err != nil {
		log.Fatal("Error opening top domains CSV file:", err)
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	reader.Comma = ','
	reader.FieldsPerRecord = -1

	rawData, err := reader.ReadAll()
	if err != nil {
		log.Fatal("Error reading CSV data:", err)
	}

	outputFile, err := os.Create("top_domains.csv")
	if err != nil {
		log.Fatal("Error creating output file:", err)
	}
	defer outputFile.Close()

	writer := csv.NewWriter(outputFile)
	defer writer.Flush()

	// Write the header row to the output CSV
	writer.Write([]string{"td", "tld", "dl", "nos", "nod", "noh", "m"})

	// Skip the header row from the input CSV
	for _, row := range rawData[1:] {
		domain := row[2]
		tld := row[3]

		domainLabels := strings.Split(domain, ".")
		domainLength := len(domainLabels)

		numSubdomains := domainLength - 1

		hyphens, numerics := countHyphensAndNumerics(domain)

		// Create a row with the preprocessed data
		preprocessedRow := []string{
			domain,
			tld,
			strconv.Itoa(domainLength),
			strconv.Itoa(numSubdomains),
			strconv.Itoa(hyphens),
			strconv.Itoa(numerics),
			"false",
		}

		// Write the preprocessed row to the output CSV
		writer.Write(preprocessedRow)
	}
}

func countHyphensAndNumerics(s string) (int, int) {
	hyphens := 0
	numerics := 0

	for _, r := range s {
		if r == '-' {
			hyphens++
		} else if unicode.IsNumber(r) {
			numerics++
		}
	}

	return hyphens, numerics
}
