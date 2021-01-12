package main

import "fmt"

func main() {
	fmt.Println(isSubsequence("axc", "ahbgdc"))
}

func isSubsequence(s string, t string) bool {
	if s == "" {
		return true
	}
	for i := 0; i < len(s) && len(t) > 0; {
		if s[i] == t[0] {
			if i == len(s)-1 {
				return true
			}
			i++
		}
		t = t[1:]
	}
	return false
}
