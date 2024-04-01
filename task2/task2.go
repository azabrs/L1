package main 

import (
	"fmt"
	"sync"
	)

func main(){
	l := []int{2, 4, 6, 8, 10}//create array 
	var wg sync.WaitGroup// WaitGroup help us synchronize? gotoutine
	for _, val := range(l){
		wg.Add(1)//add  1  member to group
		go func(val int){
			defer wg.Done() // finish 1 member in group before exit from function
			fmt.Println(val*val)
		}(val)
	}
	wg.Wait()//function wait while all count wg.Done() work do not equal wg.Add()
}
