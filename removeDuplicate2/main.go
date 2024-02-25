package main

import "fmt"

func removeDuplicates(nums []int) int {
	//to use two pointer, let's take to variable i and j
	i, j := 0, 0

	// create a temp has map

	temp := make(map[int]int)

	for i < len(nums) {

		if temp[nums[i]] < 2 {

			temp[nums[i]]++
			nums[j] = nums[i]
			j++
		}
		i++
	}
	return j
}

func main() {

	nums := []int{1, 2, 3, 3, 3, 3, 4, 4, 5, 5, 6, 7, 7, 8, 9, 9, 9, 9}

	n := removeDuplicates(nums)

	fmt.Println("The number of elements are: ", n)

}
