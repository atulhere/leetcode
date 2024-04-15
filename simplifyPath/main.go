// Given a string path, which is an absolute path (starting with a slash '/') 
// to a file or directory in a Unix-style file system, 
//  convert it to the simplified canonical path.
// In a Unix-style file system, a period '.' refers to the current directory,
// a double period '..' refers to the directory up a level, 
// and any multiple consecutive slashes (i.e. '//') are treated as a single slash '/'. 
// For this problem, any other format of periods such as '...' are treated as file/directory names.

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

func main(){

	str := "/..."  

	str = simplifyPath(str)

	fmt.Println(str)

}

func simplifyPath(str string) string{

	s :=Stack{}

	word :=""; dot :=""; output := ""

	for _, v := range str{

		
		currentChar := string(v)
		 
		if (currentChar == "/" || currentChar =="." ){

			if(word != ""){
				s.push(word)
				word =""
			}
			

			if(currentChar == "."){
				dot = dot + "."
			}
			if(dot == ".." && currentChar == "/"){
				s.pop()
				dot =""
			}

			if(dot != "" && dot !="." && dot !=".."){
				s.push(dot)
				dot =""
			}
			
			continue
		}else{
			word = word + currentChar

			dot = ""
		}

	}
	fmt.Println(s.item)
	for _,v := range s.item{
		output = output + "/"+ v.(string) 
	}
	
	if(len(s.item) == 0){
		return "/"

	}
	return output
}