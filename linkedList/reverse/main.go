package main

import "fmt"
type ListNode struct{
	Value int
	Next *ListNode
}

func reverse(head *ListNode) *ListNode {
	var prev *ListNode
	current := head 


	for current !=nil{
		
		temp := current.Next
		current.Next =prev
		prev =current
		current = temp
		
	}
	return prev
}

func display(head *ListNode){

	 current :=head

	
	for current !=nil{

		fmt.Println(current.Value)

		current = current.Next 	
	}


}

func main(){

	// let's create forward linklist, the source linkedliste 
	// that we need to reverse

	head := &ListNode{1, nil}
	head.Next = &ListNode{2,nil}
	head.Next.Next = &ListNode{3,nil}

	display(head)
	reverseList := reverse(head)
	display(reverseList)
}