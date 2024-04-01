package main

import "fmt"

func isUniq(str string) bool{
	dic := make(map[rune]bool)
	for _, val := range str{
		dic[val] = true
	}
	return len(str) == len(dic)
}

func main(){
	s1 := "abcd"
	s2 := "abCdefAaf"
	s3 := "aabcd"
	fmt.Println(isUniq(s1))
	fmt.Println(isUniq(s2))
	fmt.Println(isUniq(s3))
}