package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

type stack []string

func (s stack) Push(str string) stack {
	return append(s, str)
}

func (s stack) Pop() (stack, string) {
	l := len(s)
	return s[:l-1], s[l-1]
}

func (s stack) Tail() string {
	return s[len(s)-1]
}

func (s stack) Empty() bool {
	return len(s) == 0
}

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanRunes)

	polymers := map[string]bool{}
	sizes := map[string]int{}
	var chars []string

	for scanner.Scan() {
		c := strings.TrimRight(scanner.Text(), "\n")

		if len(c) == 0 {
			continue
		}

		polymers[strings.ToLower(c)] = true
		chars = append(chars, c)
	}

	for p := range polymers {
		s := stack{}
		for _, c := range chars {
			if strings.ToLower(c) == p {
				continue
			}

			if s.Empty() {
				s = s.Push(c)
				continue
			}

			tail := s.Tail()
			if tail != c && strings.EqualFold(c, tail) {
				s, _ = s.Pop()
				continue
			}
			s = s.Push(c)
		}
		sizes[p] = len(s)
	}

	min := math.MaxInt32

	for _, s := range sizes {
		if s < min {
			min = s
		}
	}

	fmt.Println(min)
}
