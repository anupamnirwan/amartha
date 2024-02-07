package main

import (
	"fmt"
)

func expand(word string, left, right int) string { // string
	n := len(word)
	for left >= 0 && right < n && word[left] == word[right] { //checking condition
		left--
		right++
	}
	s := word[left+1 : right] // I ran the same statement but I guess some code was unchanged left should left + 1 because we are looking for the previous index
  // and right should be right - 1 but since we are including the previous right index so right -1 +1 (to get the slice)
	return s
}

func longestPalindromicSubstring(word string) string {
	n := len(word)
	//i := 0
	ans := ""
	for i := 0; i < n; i++ {
		s := expand(word, i, i) // odd length palindromes
		if len(s) > len(ans) {
			ans = s
		}
	}

	for i := 0; i < n; i++ {
		s := expand(word, i, i+1) // even length palindromes
		if len(s) > len(ans) {
			ans = s
		}
	}

	return ans
}

func max(a, b int) bool {
	if a > b {
		return true
	}

	return false
}

func main() {
	fmt.Println(longestPalindromicSubstring("cabad"))
	fmt.Println(longestPalindromicSubstring("detyzyt"))
	fmt.Println(longestPalindromicSubstring("tattarrattat"))
}
