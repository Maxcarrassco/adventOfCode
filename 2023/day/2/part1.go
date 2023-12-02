package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)


func isValid(s string) bool {
	config := map[string]int {
		"red": 12,
		"green": 13,
		"blue": 14,
	}

	s = strings.ReplaceAll(s, ";", ",")
	res := strings.Split(s, ", ")

	for _, val := range res {
		vals := strings.Split(val, " ")
		n, _ := strconv.Atoi(vals[0])
		if n > config[vals[1]] {
			return false
		}
	}

	return true
}


func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Expected: go run <program_name>.go <test_filename>")
		return
	}
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return
	}
	sum := 0
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		vals := strings.Split(line, ": ")
		if len(vals) < 2 {
			continue
		}
		game := strings.Split(vals[0], " ")
		gameN, _ := strconv.Atoi(game[1])
		if isValid(vals[1]) {
			sum += gameN
		}
	}

	fmt.Println(sum)

}
