package main

import "fmt"


func main(){
	ch := make(chan int)
	defer close(ch)
	l := []int{2, 4, 6, 8, 10}
	go func(){
		for _, val := range(l){
			ch <- val * val
		}
	}()
	
	func(){
		sum := 0
		for i := 0; i < 5; i++{
			sum += <-ch
		}
		fmt.Println(sum)
	}()
}
