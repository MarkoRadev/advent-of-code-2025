package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

const K = 12

func maxJoltage(bank string) int {
	maxRight := -1
	largest := -1

	for i := len(bank) - 1; i >= 0; i-- {
		ch := bank[i]
		if ch < '0' || ch > '9' {
			continue
		}
		digit := int(ch - '0')

		if maxRight != -1 {
			value := 10*digit + maxRight
			if value > largest {
				largest = value
			}
		}

		if digit > maxRight {
			maxRight = digit
		}
	}

	if largest == -1 {
		return 0
	}
	return largest
}

func maxJoltageK(bank string, k int) string {
	digits := make([]byte, 0, len(bank))
	for i := 0; i < len(bank); i++ {
		ch := bank[i]
		digits = append(digits, ch)
	}

	if len(digits) <= k {
		return string(digits)
	}

	remove := len(digits) - k
	stack := make([]byte, 0, len(digits))

	for _, d := range digits {
		for remove > 0 && len(stack) > 0 && stack[len(stack)-1] < d {
			stack = stack[:len(stack)-1]
			remove--
		}
		stack = append(stack, d)
	}

	if remove > 0 {
		stack = stack[:len(stack)-remove]
	}

	if len(stack) > k {
		stack = stack[:k]
	}

	return string(stack)
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Buffer(make([]byte, 1024), 1024*1024)

	joltageSum := 0
	totalJoltage := new(big.Int)

	for scanner.Scan() {
		bank := scanner.Text()
		joltageSum += maxJoltage(bank)

		bestStr := maxJoltageK(bank, K)
		joltage := new(big.Int)
		joltage.SetString(bestStr, 10)
		totalJoltage.Add(totalJoltage, joltage)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// fmt.Println("Total output joltage:", joltageSum)
	fmt.Println("New total output joltage:", totalJoltage.String())
}
