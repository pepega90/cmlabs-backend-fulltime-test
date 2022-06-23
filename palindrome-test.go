package main

import "fmt"

func palindrome(s string) bool {
	left := 0
	right := len(s) - 1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func main() {
	// SINGLE WORD
	res1, res2, res3, res4 := palindrome("SAIPPUAKIVIKAUPPIAS"), palindrome("Aibohphobia"), palindrome("Anna"), palindrome("Civic")
	fmt.Println("SINGLE WORD")
	fmt.Println(res1, res2, res3, res4)
	fmt.Println("============")
	// MULTIPLE WORD
	fmt.Println("MULTIPLE WORD")
	resM1, resM2 := palindrome("My gym"), palindrome("No lemon, no melon")
	fmt.Println(resM1, resM2)
}
