package main

import "fmt"


type MinStack struct {
    stack []int
	minStack []int 
}


func Constructor() MinStack {
    
	return MinStack{}
}


func (this *MinStack) Push(val int)  {

	this.stack = append(this.stack, val)
	lMinStack := len(this.minStack)
	if(lMinStack == 0){
		this.minStack = append(this.minStack, val)
	}
	if(lMinStack >0){
		topMinStack := this.minStack[lMinStack -1 ]
		if(val <= topMinStack ){
			this.minStack = append(this.minStack, val)
		}
	}		   
}


func (this *MinStack) Pop()  {

	lStack := len(this.stack)
	lMinStack := len(this.minStack)

	if(lStack >0){

		//Fecth the Top element from the main stack 
		topStack := this.stack[lStack-1]
		topMinStack := this.minStack[lMinStack -1 ]

		// if both are same, pop(delete) the element from minStack too
		if(topStack == topMinStack){
			this.stack = this.stack[: lStack -1]
			this.minStack = this.minStack[: lMinStack -1]
		}else{
			this.stack = this.stack[: lStack -1]
		}

	}
  
}


func (this *MinStack) Top() int {
	l :=len(this.stack)
	if(l > 0){
		return this.stack[l-1]
	}
 return -1 

}


func (this *MinStack) GetMin() int {

	l := len(this.minStack)
	if(l > 0){
		return this.minStack[len(this.minStack) - 1]
	}
	return 0  
}

func main(){

 obj := Constructor()
 obj.Push(30)
 obj.Push(50)
 obj.Push(70)
 obj.Push(10)
 obj.Pop()
 top := obj.Top()
 min := obj.GetMin()

 fmt.Println("Top", top)
 fmt.Println("Min", min)
 

}
