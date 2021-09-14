package testmath

import (
	"fmt"
	"math"
)

func Triangle() {
	var a, b int = 3, 4
	fmt.Println("##############################", CalcTriangle(a, b))
}

func CalcTriangle(a, b int) int {
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	return c
}

func lengthOfLongestSubstring(str string) int {
	str_len := len(str)
	arr := make([]bool, 1024)
	max := 1
	count := 0
	// If string length is zero
	if str_len == 0 {
		max = 0
	} else {
		for i := 0; i < str_len; i++ {
			for j := 0; j < 125; j++ {
				arr[j] = false
			}
			count = 0
			arr[str[i]] = true
			count = 1

			for k := i + 1; k < str_len; k++ {
				if arr[str[k]] {
					break
				}
				arr[str[k]] = true
				count = count + 1
				if max < count {
					max = count
				}
			}
		}
	}
	return max
}
