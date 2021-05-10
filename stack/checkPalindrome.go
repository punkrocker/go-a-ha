package main

import "fmt"

const checkMark = "\u2713"
const errorMark = "\u2717"

func main() {
	check(false, isPalindrome("test"), "test should be false")
	check(true, isPalindrome("aba"), "aba should be true")
	check(true, isPalindrome("aa"), "aa should be true")
}

func check(want bool, real bool, message string) {
	if want == real {
		fmt.Println(message, checkMark)
	} else {
		fmt.Println(message, errorMark)
	}
}

func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}

	}
	return true
}
