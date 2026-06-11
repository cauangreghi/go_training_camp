// The challenge is found in a array two slices what them sum of twice match a target
package main

import "fmt"

func twoSum(nums []int, target int) []int{
	m := make(map[int]int)

	for idx,val := range nums{
		complement := target - val
		if compIdx, exists := m[complement];exists{
			return []int{compIdx,idx}
		}
		m[val] = idx
	}
	return []int{}
}

func main(){
	nums := []int {1,5,7,2}
	var result = twoSum(nums, 9)
	fmt.Println(result)
}