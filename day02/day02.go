package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isInvalidID(n int64) bool {
	s := strconv.FormatInt(n, 10)
	if len(s)%2 != 0 {
		return false
	}
	half := len(s) / 2
	return s[:half] == s[half:]
}

func isInvalidID_2(n int64) bool {
	s := strconv.FormatInt(n, 10)
	L := len(s)

	for subLen := 1; subLen <= L/2; subLen++ {
		if L%subLen != 0 {
			continue
		}

		repeats := L / subLen
		if repeats < 2 {
			continue
		}

		pattern := s[:subLen]
		if strings.Repeat(pattern, repeats) == s {
			return true
		}
	}

	return false
}

func main() {
	data, _ := os.ReadFile("input.txt")
	parts := strings.Split(strings.TrimSpace(string(data)), ",")

	var sum int64

	for _, p := range parts {
		if p == "" {
			continue
		}

		bounds := strings.Split(p, "-")
		start, _ := strconv.ParseInt(bounds[0], 10, 64)
		end, _ := strconv.ParseInt(bounds[1], 10, 64)

		for n := start; n <= end; n++ {
			if isInvalidID_2(n) {
				sum += n
			}
		}
	}

	fmt.Println(sum)
}
