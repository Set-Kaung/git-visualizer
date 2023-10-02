package main

import (
	"flag"
	"fmt"
)

func scan(folder string) {
	fmt.Printf("Found folder: \n\n")
	repositories := recursiveScan(folder)
	filePath := getDotFilePath()
	addNewSliceElementsToFile(filePath, repositories)
	fmt.Printf("\n\nSuccessfully added\n\n")
}

func recursiveScan(folder string) []string {
	return scanGitFolders(make([]string, 0), folder)
}

func stats(email string) {
	commits := processRepositories(email)
	printCommitStats(commits)
}

func main() {
	var folder string
	var email string
	flag.StringVar(&folder, "add", "", "add a new folder to scan for Git repos")
	flag.StringVar(&email, "email", "your@email.com", "the email to scan")
	flag.Parse()

	if folder != "" {
		scan(folder)
		return
	}
	stats(email)
}
