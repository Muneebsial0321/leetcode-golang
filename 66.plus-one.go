/*
 * @lc app=leetcode id=66 lang=golang
 *
 * [66] Plus One
 */

// @lc code=start
package main

// import (
// 	"fmt"
// 	"math/big"
// 	"strconv"
// 	"strings"
// )

// func main() {
// 	plusOne([]int{7, 2, 8, 5, 0, 9, 1, 2, 9, 5, 3, 6, 6, 7, 3, 2, 8, 4, 3, 7, 9, 5, 7, 7, 4, 7, 4, 9, 4, 7, 0, 1, 1, 1, 7, 4, 0, 0, 6})
// }
func plusOne(digits []int) []int {
	str := ""
	for _, e := range digits {
		str = str + strconv.Itoa(e)
	}
	fmt.Println("str", str)

	// fault here
	num1 := new(big.Int)
	num1.SetString(str, 10)
	// num, _ := strconv.ParseInt(str, 10, 64)
	num2:= new(big.Int)
	num2.SetString("1",10)

	ans:= new(big.Int)
	ans.Add(num1,num2)
	fmt.Println("num",ans)
	// fault here

	
	str = ans.String()
	strArray := strings.Split(str, "")
	result := []int{}
	for _, e := range strArray {
		n, _ := strconv.ParseInt(e, 10, 64)
		result = append(result, int(n))
	}
	fmt.Println("result",result)
	return result
	// return []int{}
}

// @lc code=end
