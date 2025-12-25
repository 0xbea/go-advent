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
		max_10 := 0
		max_1 := 0
		var j_start int
		for i := 0; i < length-1; i++ {
			num, err := strconv.Atoi(string(line[i]))
			if err != nil {
				fmt.Println("Error converting string to int:", err)
			}
			if num > max_10 {
				max_10 = num
				j_start = i
				if max_10 == 9 {
					break
				}
			}
		}
		for j := j_start + 1; j < length; j++ {
			num, err := strconv.Atoi(string(line[j]))
			if err != nil {
				fmt.Println("Error converting string to int:", err)
			}
			if num > max_1 {
				max_1 = num
				if max_1 == 9 {
					break
				}
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
