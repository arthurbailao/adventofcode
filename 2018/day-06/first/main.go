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

func findInfiniteAreas(grid [][]int, width, height int) map[int]bool {
	result := map[int]bool{}
	for x := 0; x < width; x++ {
		result[grid[x][0]] = true
		result[grid[x][height-1]] = true
	}

	for y := 0; y < height; y++ {
		result[grid[0][y]] = true
		result[grid[width-1][y]] = true
	}

	return result
}

func sumAreas(grid [][]int) map[int]int {
	result := map[int]int{}
	for x := range grid {
		for y := range grid[x] {
			result[grid[x][y]]++
		}
	}
	return result
}

func largestArea(areas map[int]int, infiniteAreas map[int]bool) int {
	maxArea := 0
	for i, area := range areas {
		if i < 0 {
			continue
		}

		if area > maxArea && !infiniteAreas[i] {
			maxArea = area
		}
	}

	return maxArea
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

	grid := makeGrid(width, height)

	for x := range grid {
		for y := range grid[x] {
			minDistance := math.MaxInt32

			for i, c := range coordinates {
				d := distance(xy{x, y}, c)
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

	infiniteAreas := findInfiniteAreas(grid, width, height)
	areas := sumAreas(grid)

	area := largestArea(areas, infiniteAreas)

	fmt.Printf("%d\n", area)
}
