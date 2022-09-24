// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 73.

// Comma prints its argument numbers with a comma at each power of 1000.
//
// Example:
//
//	$ go build gopl.io/ch3/comma
//	$ ./comma 1 12 123 1234 1234567890
//	1
//	12
//	123
//	1,234
//	1,234,567,890
package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("comma recursion\t%s\n", comma(os.Args[i]))
		fmt.Printf("comma loop\t%s\n", comma_loop(os.Args[i]))
	}
	s_a, s_b := "abcd", "dcae"
	fmt.Printf("str_a: %s str_b: %s similar: %t\n", s_a, s_b, isSimilarString(s_a, s_b))
}

// !+
// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	n := len(s)
	if n <= 3 || (n == 4 && (s[0] == '+' || s[0] == '-')) {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func comma_loop(s string) string {
	var buf bytes.Buffer
	n, cnt := len(s), 0
	if n <= 3 {
		return s
	}
	for i := n; i > -1; i -= 3 {
		first_ch := s[max(0, i-3)]
		if i-3 <= 0 {
			break
		}
		if first_ch == '+' || first_ch == '-' {
			break
		}
		if i-4 == 0 && (s[i-4] == '+' || s[i-4] == '-') {
			break
		}
		cnt += 1
	}
	buf.WriteString(s[:n-cnt*3])
	for i := n - cnt*3; i < n; i += 3 {
		buf.WriteString(fmt.Sprintf(",%s", s[i:i+3]))
	}
	return buf.String()
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// practice3.12
// Return true if two string have same character, not care about sequence
func isSimilarString(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	memo := [255]byte{}
	bytes_a, bytes_b := []byte(a), []byte(b)
	for i := range bytes_a {
		memo[bytes_a[i]] += 1
	}
	for i := range bytes_b {
		if memo[bytes_b[i]] > 0 {
			memo[bytes_b[i]] -= 1
		} else {
			return false
		}
	}
	return true
}

//!-
