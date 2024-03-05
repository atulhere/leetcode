package main

import "fmt"

func main() {

	nums := []int{1, 2, 3, 4, 5, 6, 7}

	fmt.Println("The given array is ", nums)

	k := 3 //rotate by 3 steps
	rotateElement(nums, k)

}

func rotateElement(nums []int, k int) {

	// Handle the case where k is greater than the length of the array
	k = k % len(nums)
	// let's revere of whole array first
	array := reverse(nums, 0, len(nums)-1)
	// reverse from start to less than k
	array = reverse(array, 0, k-1)
	// now reverse from k to end of the array
	reverse(array, k, len(array)-1)

}

/*
func rotateElementWithTemp(nums []int, k int) {
	temp := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {

		temp[(i+k)%len(nums)] = nums[i]
	}
	nums = temp

	fmt.Println("The Value of the array is ", nums)

}
*/

func reverse(nums []int, start int, end int) []int {

	i := start
	j := end

	middleElement := (i + j + 1) / 2

	for i < middleElement && j >= middleElement {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
	return nums
}
