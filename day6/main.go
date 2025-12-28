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
	// create slice of int slices

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// process each line
		// add ints to slice
		// if symbol, do operation for slice

	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
	fmt.Println("Done")
}
