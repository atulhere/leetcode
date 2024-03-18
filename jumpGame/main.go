/*
Jump Game I
You are given an integer array nums.
You are initially positioned at the array's first index,
and each element in the array represents your maximum jump
length at that position.
Return true if you can reach the last index, or false otherwise.
*/

package main

import "fmt"

func main() {

	nums := []int{3, 2, 1, 0, 4}

	jump := canJump(nums)

	fmt.Println(jump)
}
func canJump(nums []int) bool {

	destination := len(nums) - 1

	if len(nums) == 1 && nums[0] == 0 {

		return true
	}

	for i := len(nums) - 2; i >= 0; i-- {

		for j := 1; j <= nums[i]; j++ {

			if i+j == destination {
				destination = i
			}
		}
	}

	return destination == 0

}
