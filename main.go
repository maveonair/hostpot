package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"

	"github.com/maveonair/hotspot/hotspot"
)

func main() {
	cmdName := "git"
	cmdArgs := []string{"log", "--numstat", "--oneline", "--pretty=format:''"}

	output, err := exec.Command(cmdName, cmdArgs...).Output()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error executing git log command", err)
		os.Exit(1)
	}

	r := hotspot.Analyze(string(output))
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
