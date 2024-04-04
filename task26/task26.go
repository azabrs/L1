package main

import "fmt"

func isUniq(str string) bool{
	dic := make(map[rune]bool)
	for _, val := range str{//iterate over rune
		dic[val] = true
	}// create set
	return len(str) == len(dic)// compare len of set and len of string
}

func main(){
	s1 := "abcd"
	s2 := "abCdefAaf"
	s3 := "aabcd"
	fmt.Println(isUniq(s1))
	fmt.Println(isUniq(s2))
	fmt.Println(isUniq(s3))
}