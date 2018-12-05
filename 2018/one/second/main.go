package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var set map[int]bool
var total int

func findNumber(scanner *bufio.Scanner) (int, bool) {
	for scanner.Scan() {
		line := scanner.Text()
		number, _ := strconv.Atoi(line)
		total = total + number

		if set[total] {
			return total, true
		}

		set[total] = true
	}

	return total, false
}

func main() {
	set = make(map[int]bool)
	set[0] = true
	total = 0

	for {
		file, err := os.Open("../input.txt")
		defer file.Close()
		if err != nil {
			panic(err)
		}
		scanner := bufio.NewScanner(file)
		number, found := findNumber(scanner)
		if found {
			fmt.Println(number)
			break
		}

		if err = scanner.Err(); err != nil {
			panic(err)
		}
	}
}
