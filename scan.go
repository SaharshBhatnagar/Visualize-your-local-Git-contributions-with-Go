package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func scan(folder string) {

	repos := scanGitFolders(folder)
	fmt.Printf("\nFound %d repositories:\n", len(repos))

	for _, repo := range repos {
		fmt.Printf(" - %s\n", repo)
	}

	saveReposToFile(repos)
	
}

func scanGitFolders(folder string) []string {
	var repos []string

	entries, err := os.ReadDir(folder)

	if err != nil {
		return repos
	}

	for _, entry := range entries {

		if entry.IsDir() {
			dirname := entry.Name()

			if dirname == ".git" {
				repos = append(repos, folder)

				continue
			}

			if dirname == "node_modules" || dirname == "vendor" {
				continue
			}

			subfolder := filepath.Join(folder, entry.Name())
			subRepos := scanGitFolders(subfolder)
			repos = append(repos, subRepos...)

		}
	}

	return repos
}

func saveReposToFile(repos []string) {
	content := strings.Join(repos, "\n")

	err := os.WriteFile(".local_git_repos", []byte(content), 0644)

	if err != nil {
		panic(err)
	}
}