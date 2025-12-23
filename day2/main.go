package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("test_input")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// set initial / global variables
	var invalidList []map[int]bool

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// process each line ... there is only one
		idRanges := strings.Split(line, ",") // 11-22
		for _, idRange := range idRanges {
			startEnd := strings.Split(idRange, "-")
			start, err := strconv.Atoi(startEnd[0])
			if err != nil {
				print(err)
			}
			end, err := strconv.Atoi(startEnd[1])
			if err != nil {
				print(err)
			}
			// fmt.Println(start, end)
			// todo: dont bother with range where both have odd digits and same number
			// invalid ID:  any ID which is made only of some sequence of digits repeated twice.
			// generate invalid IDs for a range
			// each range has a set
			s := map[int]bool{}
			for i := start; i <= end; i++ {
				fmt.Println("i", i)
				// dont bother with IDs with odd number digits
				numDigits := len(strconv.Itoa(i))
				if numDigits%2 != 0 {
					continue
				}
				// if even number of digits, take first half and duplicate
				var firstHalf string
				var duplicated string
				firstHalf = strconv.Itoa(i)[:numDigits/2]
				duplicated = firstHalf + firstHalf
				checkId, err := strconv.Atoi(duplicated)
				if err != nil {
					fmt.Println("Error converting to int:", err)
					continue
				}
				fmt.Println("checking:", checkId)
				fmt.Println("numDigits", numDigits)
				if checkId < start {
					continue
				}
				if checkId > end {
					break
				}
				s[checkId] = true
			}
			invalidList = append(invalidList, s)
		}

	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
	fmt.Println(invalidList)
	// sum the values of the maps now
	total := 0
	for _, m := range invalidList {
		for num, _ := range m {
			total += num
		}
	}
	fmt.Println("Total:", total)
	fmt.Println("Done")
}
