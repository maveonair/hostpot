package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"sort"
	"strings"
)

func main() {
	cmdName := "git"
	cmdArgs := []string{"log", "--numstat", "--oneline", "--pretty=format:''"}

	output, err := exec.Command(cmdName, cmdArgs...).Output()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error executing git log command", err)
		os.Exit(1)
	}

	r := buildStatistic(string(output))
	printRepositoryFiles(r)
}

func buildStatistic(log string) RepositoryFiles {
	data := make(map[string]int)
	i := 0

	for _, line := range strings.Split(log, "\n") {
		splits := strings.Split(line, "\t")
		filePath := splits[len(splits)-1]

		if len(strings.TrimSpace(filePath)) != 0 && filePath != "''" {
			if _, err := os.Stat(filePath); err == nil {
				data[filePath] += 1
				i++
			}
		}
	}

	return buildRepositoryFiles(data)
}

func buildRepositoryFiles(data map[string]int) RepositoryFiles {
	repositoryFiles := make(RepositoryFiles, len(data))
	i := 0

	for filePath, churn := range data {
		lineNumber := readLineNumber(filePath)
		repositoryFiles[i] = RepositoryFile{filePath, churn, lineNumber}
		i++
	}

	sort.Sort(sort.Reverse(repositoryFiles))
	return repositoryFiles
}

func readLineNumber(filePath string) int {
	lineNumber := 0
	file, err := os.Open(filePath)
	defer file.Close()

	if err != nil {
		fmt.Fprintln(os.Stderr, "Error opening file", err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		if len(strings.TrimSpace(scanner.Text())) != 0 {
			lineNumber++
		}
	}

	return lineNumber
}

func printRepositoryFiles(repositoryFiles RepositoryFiles) {
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
