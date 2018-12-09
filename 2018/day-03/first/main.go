package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type point struct {
	x, y int
}

func extractValues(text string) [5]int {
	re := regexp.MustCompile("[0-9]+")
	ss := re.FindAllString(text, -1)
	var result [5]int
	for i, str := range ss {
		value, _ := strconv.Atoi(str)
		result[i] = value
	}

	return result
}

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fabric := map[point]int{}
	total := 0

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		values := extractValues(scanner.Text())

		x := values[1] + 1
		y := values[2] + 1
		width := values[3]
		height := values[4]

		for i := x; i < x+width; i++ {
			for j := y; j < y+height; j++ {
				p := point{i, j}
				fabric[p]++
				if fabric[p] == 2 {
					total++
				}
			}
		}

	}

	fmt.Println(total)

}
