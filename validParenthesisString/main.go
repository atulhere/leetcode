package main

import (
	"fmt"
	"errors"
)

type Stack struct{
	item []interface{}
}

func (s *Stack) push(str interface{}){

	s.item = append(s.item,str)

}


func (s *Stack) pop(){
	l :=len(s.item)
	if l > 0{
		s.item = s.item[:l-1]
	}
}

func (s *Stack) top() (interface{}, error) {

	l := len(s.item)
	if l ==0 {

		return nil, errors.New("Stack is empty")
	}
	return s.item[l-1], nil
}

func (s *Stack) isEmpty()bool{

	return len(s.item)==0

}


func validParanthesis(str string) bool{

	 s :=Stack{}

	 for i :=0; i < len(str) ; i++{
		 
		// fetch top element 
		top , err :=s.top()

		// check if current char is closing char 
		// check the character present on the top 
		// It top char is open char, pop the element present at top
		if(err !=nil) {
			s.push(string(str[i]))
			continue
		}
		if(top == "(" && string(str[i]) == ")" || top == "{" && string(str[i])== "}" || top == "[" && string(str[i])== "]"){
			s.pop()
			continue
		}
		s.push(string(str[i]))

		fmt.Println(s.item)

	 }
	return s.isEmpty()
}
func main(){

	str := "(){}"
	isValidParanthesis := validParanthesis(str)

	fmt.Println(isValidParanthesis)
	
}