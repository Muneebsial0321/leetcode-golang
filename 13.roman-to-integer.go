/*
 * @lc app=leetcode id=13 lang=golang
 *
 * [13] Roman to Integer
 */

// @lc code=start
// package main

import (
	"fmt"
	"strings"
)

// func main() { romanToInt("VIII") }
func romanToInt(s string) int {
	hashmap := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	arr := strings.Split(s, "")
	sum:=0
	for _, v := range arr {
		sum = sum + hashmap[v]
	}
	return sum
}

// @lc code=end
