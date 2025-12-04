package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	grid := make([]string, 0)
	accessible := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// process each line
		// assign to slice of slices matrix
		grid = append(grid, line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	// process grid now to count accessible rolls
	for y, row := range grid {
		for x, item := range row {
			adjacents := 0
			if item == '@' {
				// check top left
				if y > 0 && x > 0 {
					if grid[y-1][x-1] == '@' {
						adjacents++
					}
				}
				// check directly above
				if y > 0 {
					if grid[y-1][x] == '@' {
						adjacents++
					}
				}
				// check top right
				if y > 0 && x < len(row)-1 {
					if grid[y-1][x+1] == '@' {
						adjacents++
					}
				}
				// check left
				if x > 0 {
					if grid[y][x-1] == '@' {
						adjacents++
					}
				}
				// check right
				if x < len(row)-1 {
					if grid[y][x+1] == '@' {
						adjacents++
					}
				}
				// check bottom left
				if y < len(grid)-1 && x > 0 {
					if grid[y+1][x-1] == '@' {
						adjacents++
					}
				}
				// check directly below
				if y < len(grid)-1 {
					if grid[y+1][x] == '@' {
						adjacents++
					}
				}
				// check bottom right
				if y < len(grid)-1 && x < len(row)-1 {
					if grid[y+1][x+1] == '@' {
						adjacents++
					}
				}
				// if less than 4 adjacents, count as accessible
				if adjacents < 4 {
					accessible++
				}
			}
		}
	}

	fmt.Println("Total accessible rolls:", accessible)
	fmt.Println("Done")
}
