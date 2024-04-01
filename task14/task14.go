package main

import "fmt"

func determineType1(val interface{}){
	fmt.Printf("type of the variable is %T\n", val)
}

func determineType2(val interface{}){
	switch val.(type){
	case string:
		fmt.Printf("type of the variable %v is string\n", val)
	case int:
		fmt.Printf("type of the variable %v is int\n", val)
	case bool:
		fmt.Printf("type of the variable %v is bool\n", val)
	case chan int:
		fmt.Println("type of the variable is chan")

	} 
}

func main(){
	ch1 := make(chan int)
	ch2 := make(chan int)

	determineType1("Hello")
	determineType1(123)
	determineType1(true)
	determineType1(ch1)

	determineType2("Hello")
	determineType2(123)
	determineType2(true)
	determineType2(ch2)
}