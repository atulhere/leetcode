/*
Given an array of integers citations where citations[i]
is the number of citations a researcher received for
their ith paper, return the researcher's h-index.

According to the definition of h-index on Wikipedia:
The h-index is defined as the maximum value of h such
that the given researcher has published at least h
papers that have each been cited at least h times.
*/
package main

func main(){

	nums :=[]int{3,0,6,1,5}
	hIndex := hIndex(nums)

}

func hIndex(nums []int) int{

	temp :=make(map[int]int)

	for i :=0;i < len(nums); i++{
		
		if(temp[nums[i]] == 0){
			temp[nums[i]]++
		}else if(temp[nums[i]] >=1 && nums[i]){

		}
		
	}


}