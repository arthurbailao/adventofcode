package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("../input.txt")
	defer file.Close()
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		number, _ := strconv.Atoi(line)
		total = total + number
	}

	fmt.Println(total)

	if err = scanner.Err(); err != nil {
		panic(err)
	}
}
