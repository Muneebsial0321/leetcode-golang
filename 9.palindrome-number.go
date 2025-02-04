/*
 * @lc app=leetcode id=9 lang=golang
 *
 * [9] Palindrome Number
 */
// package main

import (
	// "fmt"
	"strconv"
	"strings"
)

// @lc code=start
// func main() { isPalindrome(121) }
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	l := strconv.Itoa(x)
	arr := strings.Split(l, "")
	start := 0
	end := len(arr) - 1
	for i := 0; i < len(arr)-1; i++ {
		if start == end {
			return true
		}
		if arr[start] != arr[end] {
			return false
		}
		start++
		end--

	}
	return true

}

// @lc code=end
