package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("test_input")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	total_jolts := 0
	for scanner.Scan() {
		line := scanner.Text()
		// process each line
		max_10 := 0
		max_1 := 0
		for _, char := range line {
			num, _ := strconv.Atoi(string(char))
			if num > max_1 {
				max_10 = max_1
				max_1 = num
			}
		}
		joltage := max_10*10 + max_1
		fmt.Println("joltage for line:", joltage)
		total_jolts += joltage
	}
	fmt.Println("Total joltage:", total_jolts)

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
	fmt.Println("Done")
}
