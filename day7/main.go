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

	scanner := bufio.NewScanner(file)
	beamMap := [][]string
	splits := 0 
	i := 0
	// .......S.......
	// .......|.......
	// ......|^|......
	for scanner.Scan() {
		// process each line
		line := scanner.Text()
		if i == 0 {
			beamMap[0] = line
		}
		if i > 0 {
			if { // if line above contains S
				beamMap[i] = "" // add | under S  
			}
			// count the number of splitter with tachyon beam above it 
			// compare line being processed with line above 
			for idx, r := range(line) {
				if r == '^' {

				} else if beamMap[i-1][idx] == '|' {

				} else beamMap[i][idx] := '.' // set char 
			}
		} 
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
	fmt.Println("Done")
}
