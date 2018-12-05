package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("../input.txt")
	defer file.Close()
	if err != nil {
		panic(err)
	}

	checksum := map[int]int{2: 0, 3: 0}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		chrsCount := map[rune]int{}
		for _, chr := range scanner.Text() {
			chrsCount[chr]++
		}

		values := map[int]int{}
		for _, v := range chrsCount {
			values[v] = 1
		}

		checksum[2] = checksum[2] + values[2]
		checksum[3] = checksum[3] + values[3]

	}
	fmt.Println(checksum[2] * checksum[3])
}
