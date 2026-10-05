package main

import (
	"fmt"
	"os"
	"strings"
	"time"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

func parseFileLines() []string {

	bytes, err := os.ReadFile(".local_git_repos")

	if err != nil {
		fmt.Println("No repositories found. Run the scanner first.")

		panic(err)
	}

	content := string(bytes)

	result:= strings.Split(content, "\n")

	return result
}

func fillCommits(email string, path string, commits map[int]int) map[int]int {
	repo, err := git.PlainOpen(path)

	if err != nil {
		return commits
	}

	ref, err := repo.Head()

	if err != nil {
		return commits
	}

	iterator, err := repo.Log(&git.LogOptions{From: ref.Hash()})

	if err != nil {
		return commits
	}

	cutoff := time.Now().AddDate(-1, 0, 0)


	iterator.ForEach(func(c *object.Commit) error {

		if c.Author.Email != email {
			return nil
		}

		if c.Author.When.Before(cutoff) {
			return nil
		}

		midnight := time.Date(c.Author.When.Year(), c.Author.When.Month(), c.Author.When.Day(), 0, 0, 0, 0, c.Author.When.Location()).Unix()
		commits[int(midnight)]++

		return nil
	})

	return commits
}

func stats(email string) {
	repos := parseFileLines()

	commits := make(map[int]int)

	for _, path := range repos {
		commits = fillCommits(email, path, commits)
	}
	fmt.Printf("Found commits on %d unique days\n", len(commits))

	printCommits(commits)
}

