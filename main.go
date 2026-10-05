package main

import (
	"flag"
	"fmt"
)

func main() {

	folder := flag.String("folder", "", "adds a folder to scan")

	email := flag.String("email", "", "adds a email to scan")

	flag.Parse()

	if *folder != "" {
		// fmt.Println("Scanning Directory:", *folder)
		scan(*folder)

		return
	}

	if *email != "" {
		fmt.Println("Scanning email:", *email)
		return
	}
}