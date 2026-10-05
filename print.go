package main

import (
	"fmt"
	"time"
)

func buildCell(val int, maxValue int) string {

	if val == 0 {
        return "\033[38;5;237m■ \033[0m"
    }

    if val <= (maxValue / 4) {
        return "\033[38;5;22m■ \033[0m"
    }

    if val <= (maxValue / 2) {
        return "\033[38;5;28m■ \033[0m"
    }

    if val <= (maxValue - (maxValue / 4)) {
        return "\033[38;5;40m■ \033[0m"
    } else {
        return "\033[38;5;46m■ \033[0m"
    }
	
}

func printCells(commits map[int]int, maxValue int) {
	today := time.Now()

	today = time.Date(today.Year(), today.Month(), today.Day(), 0, 0, 0, 0, today.Location())

	currentWeekday := int(today.Weekday())

	for row := 0; row < 7; row++ {
		for col := 0; col < 53; col++ {
			daysAgo := (52 - col)*7 + (currentWeekday - row)

			if daysAgo < 0 || daysAgo > 365 {
				fmt.Print("  ")
				continue
			}

			cellDate := today.AddDate(0, 0, -daysAgo)

			timestamp := int(cellDate.Unix())

			count := commits[timestamp]

			color := buildCell(count, maxValue)

			fmt.Print(color)
		}
	
		fmt.Println()
	}

}

func printCommits(commits map[int]int) {
	maxValue := 0

	for _, count := range commits {
		if count > maxValue {
			maxValue = count
		}

	}

	fmt.Printf("Your higest commit count in a single day was: %d\n", maxValue)

	printCells(commits, maxValue)
}