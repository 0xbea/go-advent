package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// set initial / global variables

type coord struct {
	x int
	y int
}

var coordList []coord

func main() {
	file, err := os.Open("input")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	maxArea := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// process each line
		// make a coord
		xY := strings.Split(line, ",")
		x, _ := strconv.Atoi(xY[0])
		y, _ := strconv.Atoi(xY[1])
		newCoord := coord{
			x: x,
			y: y,
		}
		// fmt.Println(newCoord)
		coordList = append(coordList, newCoord)
		// fmt.Println(coordList)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	for i := 0; i < len(coordList)-1; i++ {
		for j := i + 1; j < len(coordList); j++ {
			corner0 := coordList[i]
			corner1 := coordList[j]
			width := corner0.x - corner1.x
			if width < 0 {
				width = -width
			}
			width += 1 // include the coords
			length := corner0.y - corner1.y
			if length < 0 {
				length = -length
			}
			length += 1
			area := width * length
			// fmt.Println(area)
			if area > maxArea {
				maxArea = area
			}
		}
	}

	fmt.Println("Done")
	fmt.Println("Maximum Area:", maxArea)
}
