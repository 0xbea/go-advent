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
	// simple, slow simulation: move one step at a time and count when landing exactly at 0
	pos := 50 // visible dial 0..99
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
		if dir == 'R' {
			for i := 0; i < turn; i++ {
				pos = (pos + 1) % 100
				if pos == 0 {
					pass++
				}
			}
		} else if dir == 'L' {
			for i := 0; i < turn; i++ {
				pos = (pos - 1 + 100) % 100
				if pos == 0 {
					pass++
				}
			}
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
	fmt.Println("Final pass count:", pass)
}
