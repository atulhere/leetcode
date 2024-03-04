package main

import "fmt"

func majorityElementOptimum(nums []int) int {

	count := 1
	index := 0

	for i := 1; i < len(nums); i++ {

		if nums[index] == nums[i] {
			count++
		} else {
			count--
		}

		if count == 0 {
			index = i
			count = 1
		}
	}

	m := 0
	fmt.Println("The majority value seems to be ", nums[index])
	for i := 0; i < len(nums); i++ {
		if nums[i] == nums[index] {
			m++
		}
	}
	if m > len(nums)/2 {
		return nums[index]
	} else {
		return -1
	}

}

func majorityElement(nums []int) int {

	m := len(nums) / 2
	i := 0
	temp := make(map[int]int)
	var n int

	for i < len(nums) {
		temp[nums[i]]++
		if temp[nums[i]] > m {
			n = nums[i]
		}
		i++
	}

	return n

}
func main() {
	nums := []int{1, 2, 3, 4, 4, 4, 4, 4, 4, 5, 5}
	m := majorityElementOptimum(nums)

	fmt.Println("The major Element is ", m)

}

// Another approach
// Sort the array and take the middle element of the array
