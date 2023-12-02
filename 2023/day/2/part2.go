package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)


func getPowerOfSet(s string) int {
	scoreMap := map[string]int{}
	prod := 1

	s = strings.ReplaceAll(s, ";", ",")
	res := strings.Split(s, ", ")

	for _, val := range res {
		vals := strings.Split(val, " ")
		n, _ := strconv.Atoi(vals[0])
		if n > scoreMap[vals[1]] {
			scoreMap[vals[1]] = n
		}
	}
	for _, v := range scoreMap {
		prod *= v
	}
	return prod
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
		sum += getPowerOfSet(vals[1])
	}

	fmt.Println(sum)

}
