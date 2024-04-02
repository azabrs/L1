package main

import (
	"math/rand"

)

var justString string

func createHugeString(size int) string {
	letter := "abcdefghijklmnopqrstuvwxyz"
	runes := []rune(letter)
	buffer := make([]rune, size)
	for i := range buffer{
		buffer[i] = runes[rand.Intn(len(runes))]
	}

	return string(buffer)
}


func someFunc() {
  v := createHugeString(1 << 10)
  justString = string(append([]rune{}, []rune(v)[:100]...))
}

func main() {
  someFunc()
}
