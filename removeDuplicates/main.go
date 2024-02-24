/*
Given an integer array nums sorted in non-decreasing order, remove the duplicates in-place such that each unique element appears only once. The relative order of the elements should be kept the same. Then return the number of unique elements in nums.

Consider the number of unique elements of nums to be k, to get accepted, you need to do the following things:

Change the array nums such that the first k elements of nums contain the unique elements in the order they were present in nums initially. The remaining elements of nums are not important as well as the size of nums.
Return k.
*/

package main

import "fmt"

func removeDuplicates(nums []int) int {

	// create a temp HASH map
	temp := make(map[int]bool)

	// Use two poiter approach

	i, j := 0, 0
	for i < len(nums) {
		if !temp[nums[i]] {

			temp[nums[i]] = true
			nums[j] = nums[i]

			j++
		}
		i++
	}
	return j
}

func main() {

	nums := []int{1, 1, 3, 10, 91, 3}
	fmt.Println("Given Array ", nums)
	n := removeDuplicates(nums)
	fmt.Println("Number of unqiue values are ", n)
}
