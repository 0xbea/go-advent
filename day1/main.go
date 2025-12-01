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
	dial := 50 // starting point of dial
	pass := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// process each line
		dir := line[0]
		turn, err := strconv.Atoi(line[1:])
		if err != nil {
			fmt.Println("Error converting string to int:", err)
			continue
		}
		// if L, subtract
		if dir == 'L' {
			dial -= turn
		}
		// if R, add
		if dir == 'R' {
			dial += turn
		}
		// Wrap dial to 0-99 range
		dial = ((dial % 100) + 100) % 100
		// if dial is exactly 0, increment pass
		if dial == 0 {
			pass++
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
	fmt.Println("Final pass count:", pass)
}
