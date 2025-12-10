package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("test_input")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// set initial / global variables

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// process each line
		var lights []bool
		// convert what's between brackets to boolean slice

	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
	fmt.Println("Done")
}
