/*
 * @lc app=leetcode id=58 lang=golang
 *
 * [58] Length of Last Word
 */

// @lc code=start
package main

// import (
// 	"fmt"
// 	"strings"
// )

// func main() { lengthOfLastWord("   fly me   to   the moon  ") }
func lengthOfLastWord(s string) int {
	str:= strings.TrimSpace(s)
	arrOfWords := strings.Split(str, " ")
	return len(arrOfWords[len(arrOfWords)-1])
}

// @lc code=end
