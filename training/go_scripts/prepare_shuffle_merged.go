package main

import (
	"encoding/csv"
	"math/rand"
	"os"
	"time"
)

func main() {
	// Read the merged_domains.csv
	mergedFile, err := os.Open("merged_domains.csv")
	checkError(err)

	mergedReader := csv.NewReader(mergedFile)
	mergedDomains, err := mergedReader.ReadAll()
	checkError(err)
	mergedFile.Close()

	// Shuffle the lines, excluding the header
	rand.Seed(time.Now().UnixNano())
	shuffledDomains := mergedDomains[1:]
	rand.Shuffle(len(shuffledDomains), func(i, j int) {
		shuffledDomains[i], shuffledDomains[j] = shuffledDomains[j], shuffledDomains[i]
	})

	// Write the shuffled data to a new CSV file
	shuffledFile, err := os.Create("shuffled_domains.csv")
	checkError(err)

	shuffledWriter := csv.NewWriter(shuffledFile)
	err = shuffledWriter.Write(mergedDomains[0]) // Write the header
	checkError(err)

	err = shuffledWriter.WriteAll(shuffledDomains)
	checkError(err)
	shuffledWriter.Flush()
	shuffledFile.Close()
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
