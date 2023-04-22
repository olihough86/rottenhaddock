package main

import (
	"encoding/csv"
	"os"
	"strings"
	"sync"
)

func main() {
	// Read the top_domains.csv
	topFile, err := os.Open("top_domains.csv")
	checkError(err)

	topReader := csv.NewReader(topFile)
	topDomains, err := topReader.ReadAll()
	checkError(err)
	topFile.Close()

	// Read the bad_domains.csv
	badFile, err := os.Open("bad_domains.csv")
	checkError(err)

	badReader := csv.NewReader(badFile)
	badDomains, err := badReader.ReadAll()
	checkError(err)
	badFile.Close()

	filteredBadDomains := filterBadDomains(topDomains, badDomains)

	// Adjust the number of lines in top_domains.csv
	topDomains = topDomains[:1+len(filteredBadDomains)-1]

	// Merge the files
	merged := append(topDomains, filteredBadDomains[1:]...)

	// Write the merged data to a new CSV file
	mergedFile, err := os.Create("merged_domains.csv")
	checkError(err)

	mergedWriter := csv.NewWriter(mergedFile)
	err = mergedWriter.WriteAll(merged)
	checkError(err)
	mergedWriter.Flush()
	mergedFile.Close()
}

func filterBadDomains(topDomains, badDomains [][]string) [][]string {
	filteredBadDomains := [][]string{badDomains[0]} // Include header
	var wg sync.WaitGroup
	badDomainChan := make(chan []string)

	for _, badDomain := range badDomains[1:] {
		wg.Add(1)
		go func(badDomain []string) {
			defer wg.Done()
			found := false
			for _, topDomain := range topDomains[1:] {
				if strings.ToLower(topDomain[0]) == strings.ToLower(badDomain[0]) {
					found = true
					break
				}
			}
			if !found {
				badDomainChan <- badDomain
			}
		}(badDomain)
	}

	go func() {
		wg.Wait()
		close(badDomainChan)
	}()

	for badDomain := range badDomainChan {
		filteredBadDomains = append(filteredBadDomains, badDomain)
	}

	return filteredBadDomains
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
