package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type xy struct {
	x, y int
}

func findLimits(coordinates []xy) (xy, xy) {
	maxX, maxY, minX, minY := coordinates[0].x, coordinates[0].y, coordinates[0].x, coordinates[0].y

	for _, c := range coordinates {
		if c.x > maxX {
			maxX = c.x
		}

		if c.x < minX {
			minX = c.x
		}

		if c.y > maxY {
			maxY = c.y
		}

		if c.y < minY {
			minY = c.y
		}
	}

	return xy{minX, minY}, xy{maxX, maxY}
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

func makeGrid(width, height int) [][]int {
	grid := make([][]int, width)
	for i := range grid {
		grid[i] = make([]int, height)
	}
	return grid
}

func printGrid(grid [][]int, width, height int) {

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			value := grid[x][y]
			if value < 0 {
				fmt.Print(".")
				continue
			}
			fmt.Printf("%s", string(value+65))
		}
		fmt.Printf("\n")
	}
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

	_, bottomRight := findLimits(coordinates)
	width := bottomRight.x + 1
	height := bottomRight.y + 1

	grid := makeGrid(width, height)

	for x := range grid {
		for y := range grid[x] {
			minDistance := math.MaxInt32

			for i, c := range coordinates {
				d := distance(xy{x, y}, c)
				fmt.Println(d)
				if d < minDistance {
					grid[x][y] = i
					minDistance = d
				} else if d == minDistance {
					grid[x][y] = -1
				}

			}
		}
	}

	printGrid(grid, width, height)

	fmt.Printf("%+v\n", coordinates)
}
