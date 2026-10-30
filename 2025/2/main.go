package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var divisors = make(map[int][]int)

func readInput(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	lines := []string{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err = scanner.Err(); err != nil {
		return nil, err
	}

	var idRanges []string
	for _, line := range lines {
		line = strings.TrimSpace(line)
		idRanges = append(idRanges, strings.Split(line, ",")...)
	}

	return idRanges, nil
}

func getInvalidIds(ranges []string) ([]int, error) {
	var invalidIds []int
	seen := make(map[string]bool)

	for _, r := range ranges {
		parts := strings.Split(r, "-")
		low, err := strconv.Atoi(parts[0])
		if err != nil {
			return []int{}, nil
		}
		high, err := strconv.Atoi(parts[1])
		if err != nil {
			return []int{}, nil
		}
		for i := low; i <= high; i++ {
			if seen[strconv.Itoa(i)] {
				continue
			}
			if !isValid2(strconv.Itoa(i)) {
				invalidIds = append(invalidIds, i)
			}
			seen[strconv.Itoa(i)] = true
		}
	}

	return invalidIds, nil
}

func isValid(val string) bool {
	n := len(val)
	if (n % 2) != 0 {
		return true
	}
	return val[:n/2] != val[n/2:]
}

func isValid2(val string) bool {
	n := len(val)
	if _, ok := divisors[n]; !ok {
		divisors[n] = getDivisors(n)
	}
	divs := divisors[n]

	for _, d := range divs {
		if d > n / 2 {
			continue
		}
		if val == strings.Repeat(val[:d], n / d) {
			return false
		}
	}

	return true
}

func getDivisors(x int) []int {
	ret := make([]int, 0)
	for i := 1; i*i <= x; i++ {
		if x%i != 0 {
			continue
		}
		if i == x/i {
			ret = append(ret, i)
		} else {
			ret = append(ret, i, x/i)
		}
	}
	return ret
}

func run() (int, error) {
	ranges, err := readInput("input.txt")
	if err != nil {
		return 0, err
	}

	ids, err := getInvalidIds(ranges)
	if err != nil {
		return 0, err
	}

	sum := 0
	for _, id := range ids {
		sum += id
	}

	return sum, nil
}

func main() {
	sum, err := run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unexpected error: %s", err)
		os.Exit(1)
	}
	fmt.Fprintf(os.Stdout, "Invalid sum: %d", sum)
}
