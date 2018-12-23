package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type xy struct {
	x, y int
}

func findBottomRight(coordinates []xy) xy {
	maxX, maxY := coordinates[0].x, coordinates[0].y

	for _, c := range coordinates {
		if c.x > maxX {
			maxX = c.x
		}

		if c.y > maxY {
			maxY = c.y
		}
	}

	return xy{maxX, maxY}
}

func distance(a xy, b xy) int {
	return abs(a.x-b.x) + abs(a.y-b.y)
}

func abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}

func main() {

	file, err := os.Open("../input.txt")
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()

	var coordinates []xy

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Split(line, ", ")
		x, _ := strconv.Atoi(values[0])
		y, _ := strconv.Atoi(values[1])

		coordinates = append(coordinates, xy{x, y})
	}

	bottomRight := findBottomRight(coordinates)
	width := bottomRight.x + 1
	height := bottomRight.y + 1

	var size int
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			var totalDistance int

			for _, c := range coordinates {
				totalDistance += distance(xy{x, y}, c)

			}

			if totalDistance < 10000 {
				size++
			}

		}
	}

	fmt.Println(size)

}
