package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	start int
	end   int
}

func isInRanges(num int, ranges []Range) bool {
	for _, r := range ranges {
		if num >= r.start && num <= r.end {
			return true
		}
	}
	return false
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var freshRanges []Range
	have_fresh := map[int]bool{}
	spoiled := map[int]bool{}

	scanner := bufio.NewScanner(file)

	// Phase 1: Read ranges of fresh ingredients
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			// Empty line marks transition to phase 2
			break
		}

		// Parse range like "3-5"
		ends := strings.Split(line, "-")
		start, _ := strconv.Atoi(ends[0])
		end, _ := strconv.Atoi(ends[1])
		freshRanges = append(freshRanges, Range{start: start, end: end})
	}

	// Phase 2: Check individual ingredients
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		ingredient_number, _ := strconv.Atoi(line)
		if isInRanges(ingredient_number, freshRanges) {
			have_fresh[ingredient_number] = true
		} else {
			spoiled[ingredient_number] = true
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
	fmt.Println("number of fresh ingredients:", len(have_fresh))
	fmt.Println("Done")
}
