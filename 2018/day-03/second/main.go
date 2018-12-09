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

type claim struct {
	id     int
	points []point
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
	var claims []claim

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		values := extractValues(scanner.Text())

		c := claim{}
		c.id = values[0]

		x := values[1] + 1
		y := values[2] + 1
		width := values[3]
		height := values[4]

		for i := x; i < x+width; i++ {
			for j := y; j < y+height; j++ {
				p := point{i, j}
				c.points = append(c.points, p)
				fabric[p]++
			}
		}

		claims = append(claims, c)

	}

	for _, c := range claims {
		found := true
		for _, p := range c.points {
			if fabric[p] > 1 {
				found = false
				break
			}
		}

		if found {
			fmt.Println(c.id)
			break
		}
	}

}
