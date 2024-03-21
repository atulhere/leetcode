//Problems from Coderbyte
//https://go.dev/play/p/KQiQqWU5CP1
package main

import "fmt"

func removeChar(str string, array []string) string {

	temp := make(map[string]bool)
	outPutstr := ""
	for i := 0; i < len(array); i++ {

		temp[array[i]] = true
	}
	for j := 0; j < len(string(str)); j++ {
		if !temp[string(str[j])] {

			outPutstr = outPutstr + string(str[j])
		}
	}

	return outPutstr

}

func main() {

	str := "hello world"

	array := []string{"h", "e", "w", "o"}

	str = removeChar(str, array)

	fmt.Println(str)
}
