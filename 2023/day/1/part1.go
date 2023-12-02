package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)


func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}


func getCaliVal(s string) int {
	var l_val byte
	var h_val byte
	if len(s) < 1 {
		return 0
	}

	for _, ch := range s {
		c := byte(ch)
		if isDigit(c) {
			l_val = c
			break
		}
	}
	for i := len(s) - 1; i >= 0; i--{
		if isDigit(s[i]) {
			h_val = s[i]
			break
		}
	}
	out, _ := strconv.Atoi(fmt.Sprintf("%v%v", l_val - '0', h_val - '0'))

	return out
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
	lines := strings.Split(string(data), "\n")
	sum := 0

	for _, line := range lines {
		val := getCaliVal(line)
		sum += val
	}
	fmt.Println(sum)
}
