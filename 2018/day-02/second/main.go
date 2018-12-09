package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

func oneCharDiff(f1 []string, f2 []string) (string, bool) {
	var buf bytes.Buffer
	diffs := 0
	for i, c1 := range f1 {
		c2 := f2[i]
		if c1 == c2 {
			buf.WriteString(c1)
		} else {
			diffs++
		}

		if diffs > 1 {
			return buf.String(), false
		}
	}

	return buf.String(), diffs == 1
}

func main() {
	file, err := os.Open("../input.txt")
	defer file.Close()
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	fingerprints := [][]string{}
	for scanner.Scan() {
		fingerprint := strings.Split(scanner.Text(), "")
		fingerprints = append(fingerprints, fingerprint)
	}

	if err = scanner.Err(); err != nil {
		panic(err)
	}

	for i := 0; i < len(fingerprints); i++ {
		for j := i + 1; j < len(fingerprints); j++ {
			f1 := fingerprints[i]
			f2 := fingerprints[j]

			result, found := oneCharDiff(f1, f2)
			if found {
				fmt.Printf("%s ~ %s = %s\n", strings.Join(f1, ""), strings.Join(f2, ""), result)
				break
			}
		}
	}
}
