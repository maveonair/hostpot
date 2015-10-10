package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/maveonair/hotspot/hotspot"
)

func main() {
	r := hotspot.Analyze()
	printRepositoryFiles(r)
}

func printRepositoryFiles(repositoryFiles hotspot.RepositoryFiles) {
	formatPtr := flag.String("format", "csv", "available formats: csv, json")
	flag.Parse()

	if *formatPtr == "json" {
		json := repositoryFiles.ToJson()
		fmt.Println(json)
		os.Exit(0)
	}

	csv := repositoryFiles.ToCSV()
	fmt.Println(csv)
}
