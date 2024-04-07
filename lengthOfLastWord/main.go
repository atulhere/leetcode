package main

import "fmt"

func main() {

	s := "luffy is still joyboy"

	l := lengthOfLastWord(s)

	fmt.Println(l)

}

func lengthOfLastWord(s string) int {

	//check if last char is having space, remove the space
	count := 0

	for i := len(s) - 1; i >= 0; i-- {

		if string(s[i]) != " " {
			count++
		}
		if string(s[i]) == " " && count >=1 {
			break
		}
	
	}
	return count
}
