package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

type entry struct {
	time time.Time
	text string
}

func readInput(path string) []string {
	file, err := os.Open(path)
	defer file.Close()

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	var input []string
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	return input
}

func parseLine(line string) entry {
	re := regexp.MustCompile("\\[(\\d+-\\d+-\\d+ \\d+:\\d+)\\] (.+)")
	matches := re.FindStringSubmatch(line)
	t, err := time.Parse("2006-01-02 15:04", matches[1])
	if err != nil {
		panic(err)
	}

	return entry{t, matches[2]}
}

func main() {
	input := readInput("../input.txt")
	var entries []entry

	for _, line := range input {
		e := parseLine(line)
		entries = append(entries, e)
	}

	sort.Slice(entries, func(i, j int) bool {
		return entries[i].time.Before(entries[j].time)
	})

	guards := map[string][60]int{}

	var id string
	var start time.Time
	re := regexp.MustCompile("\\d+")

	for _, e := range entries {
		if strings.HasPrefix(e.text, "Guard") {
			id = re.FindString(e.text)
		} else if strings.HasPrefix(e.text, "falls") {
			start = e.time
		} else if strings.HasPrefix(e.text, "wakes") {
			minutes := guards[id]

			for start.Before(e.time) {
				minutes[start.Minute()]++
				start = start.Add(time.Minute)
			}
			guards[id] = minutes
		}
	}

	var minute, max int
	var guard string
	for g, minutes := range guards {
		for m, sum := range minutes {
			if sum > max {
				guard = g
				minute = m
				max = sum
			}
		}
	}

	g, _ := strconv.Atoi(guard)
	fmt.Println(g * minute)
}
