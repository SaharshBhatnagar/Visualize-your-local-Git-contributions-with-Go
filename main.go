package main

import (
	"flag"
)

func main() {

	folder := flag.String("folder", "", "adds a folder to scan")

	email := flag.String("email", "", "adds a email to scan")

	flag.Parse()

	if *folder != "" {
		
		scan(*folder)
		return
	}

	if *email != "" {
		stats(*email)
		return
	}
}