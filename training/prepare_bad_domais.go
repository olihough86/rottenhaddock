package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	sourceFile := "raw.csv"
	outputFile := "training_data_raw_mal.csv"

	file, err := os.Open(sourceFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ';'

	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	outFile, err := os.Create(outputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()

	writer := csv.NewWriter(outFile)
	defer writer.Flush()

	headers := []string{"td", "tld", "dl", "nos", "nod", "noh", "m"}
	writer.Write(headers)

	for _, record := range records[1:] { // Skip the first row as it contains the headers
		taskDomain := record[6]
		tld := extractTLD(taskDomain)
		domainLength := len(taskDomain)
		numSubdomains := strings.Count(taskDomain, ".")
		numDigits := countDigits(taskDomain)
		numHyphens := strings.Count(taskDomain, "-")
		malicious := record[18]

		newRow := []string{
			taskDomain,
			tld,
			strconv.Itoa(domainLength),
			strconv.Itoa(numSubdomains),
			strconv.Itoa(numDigits),
			strconv.Itoa(numHyphens),
			malicious,
		}

		err := writer.Write(newRow)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Data transformation completed. The new CSV file is:", outputFile)
}

func countDigits(s string) int {
	count := 0
	for _, char := range s {
		if char >= '0' && char <= '9' {
			count++
		}
	}
	return count
}

func extractTLD(domain string) string {
	parts := strings.Split(domain, ".")
	if len(parts) > 1 {
		return parts[len(parts)-1]
	}
	return ""
}
