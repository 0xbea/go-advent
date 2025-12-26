package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	total_jolts := 0
	for scanner.Scan() {
		line := scanner.Text()
		length := len(line)
		freeSlots := 12
		var indexes [12]int
		var joltage [12]int

		// find max within the limits
		for f := 0; f < freeSlots; f++ {
			var i_start int
			if f == 0 {
				i_start = 0
			} else {
				i_start = indexes[f-1] + 1
			}
			for i := i_start; i <= length-(12-f); i++ {
				candidate, err := strconv.Atoi(string(line[i]))
				if err != nil {
					fmt.Println("Error converting string to int:", err)
				}
				if candidate > joltage[f] {
					// store candidate and index
					joltage[f] = candidate
					indexes[f] = i
					if candidate == 9 {
						break
					}
				}
			}
		}
		fmt.Println("joltage for line:", joltage)
		joltageValue := 0
		for _, digit := range joltage {
			joltageValue = joltageValue*10 + digit
		}
		total_jolts += joltageValue
	}
	fmt.Println("Total joltage:", total_jolts)

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
	fmt.Println("Done")
}
