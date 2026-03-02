package main

import "log"

func main() {

	s := []string{
		"()",
		"()[]{}",
		"(]",
		"([)]",
		"{[]}",
	}

	for _, v := range s {
		log.Print(isValid(v))
	}
}

func isValid(s string) bool {
	sr := []rune{}

	p := map[rune]rune{
		')': '(',
		']': '[',
		'}': '}',
	}

	for _, ch := range s {
		if ch == '(' || ch == '[' || ch == '{' {
			sr = append(sr, ch)
			continue
		}

		if len(sr) == 0 {
			return false
		}

		f := sr[len(sr)-1]
		if f != p[ch] {
			return false
		}

		sr = sr[:len(sr)-1]
	}

	return len(sr) == 0
}
