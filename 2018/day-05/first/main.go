package main

import (
	"bufio"
	"fmt"
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

	s := stack{}

	for scanner.Scan() {
		c := strings.TrimRight(scanner.Text(), "\n")

		if len(c) == 0 {
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

	fmt.Println(len(s))
}
