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

func getDigit(s int, src string) byte {
	dictH := map[string]int{
		"ze": 2,
		"on": 1,
		"tw": 1,
		"th": 3,
		"fo": 2,
		"fi": 2,
		"si": 1,
		"se": 3,
		"ei": 3,
		"ni": 2,
	}
	dictF := map[string] byte {
		"zero": '0',
		"one": '1',
		"two": '2',
		"three": '3',
		"four": '4',
		"five": '5',
		"six": '6',
		"seven": '7',
		"eight": '8',
		"nine": '9',
	}
	if s + 2 > len(src) {
		return 'z'
	}
	n, ok := dictH[src[s: s + 2]]
	if !ok {
		return 'z'
	}
	s_len := s + n + 2
	if s_len > len(src) {
		return 'z'
	}
	if val, ok := dictF[src[s: s_len]]; ok {
		return val
	}
	return 'z'
}

func getCaliVal(s string) int {
	var l_val byte
	var h_val byte
	if len(s) < 1 {
		return 0
	}

	for idx, ch := range s {
		if n := getDigit(idx, s); n != 'z' {
			l_val = n
			break
		}
		c := byte(ch)
		if isDigit(c) {
			l_val = c
			break
		}
	}
	for i := len(s) - 1; i >= 0; i--{
		idx := 0
		if i > 0 {
			idx = i - 1
		} else {
			idx = i
		}
		if n := getDigit(idx, s); n != 'z' {
			h_val = n
			break
		}
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
