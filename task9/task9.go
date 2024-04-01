package main

import (
	"fmt"
	"sync"
)

func pipelneBegin(in <-chan int, out chan<- int){
	for{
		if buf, opened := <-in; opened{
			out <- buf
		}else{
			break
		}
		
		
	}
	close(out)
}

func pipelneEnd(in <-chan int, s *sync.WaitGroup){
	for{
		if buf, opened := <-in; opened{
			fmt.Println(buf * buf)
		}else{
			break
		}
	}
	s.Done()
}

func main(){
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ch1 := make(chan int)
	ch2 := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	go pipelneBegin(ch1, ch2)
	go pipelneEnd(ch2, &wg)
	for _, val := range(arr){
		ch1 <- val
	}
	close(ch1)
	wg.Wait()
}