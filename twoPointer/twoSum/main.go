package main

import "fmt"
func main(){

	nums :=[]int{2,7,16,25}
	target := 9
	arr := twoSum(nums, target)
	fmt.Println("Pair for two sum is : ", arr)

}
func twoSum(nums []int, target int)[]int{

	//create a slice 
	arr :=[]int{}
	left :=0
	right := len(nums) -1 

	for (left < right){

		if (nums[left] + nums[right] == target){

			arr = append(arr, left +1 )
			arr = append(arr, right +1)
		}

		if((nums[left] + nums[right]) < target){
			left ++
		}else{
			right--
		}

		
      

	}
	return arr
}