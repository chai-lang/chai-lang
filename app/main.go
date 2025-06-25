package main

import (
	"chai-lang/app/scanner"
	"fmt"
)

func main() {
	scanner := &scanner.Scanner{}

	err := scanner.Scan()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Scan completed successfully.")
	}
}
