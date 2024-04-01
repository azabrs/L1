package main

import ("time"
	"fmt"
	"log"
	)

func worker(ch chan interface{}){
	for{
		buf := <-ch
		fmt.Println(buf)
	}
}

func main(){
	fmt.Println("Input number of worker")
	var n int
	_, err := fmt.Scan(&n)
	if err != nil{
		log.Fatal(err)
	}
	ch := make(chan interface{}, n)
	defer close(ch)
	for i := 0; i < n; i ++{
		go worker(ch)
	}
	var buffer string
	for{
		fmt.Println("Input data for send to worker")
		fmt.Scan(&buffer)
		fmt.Println()
		for i := 0; i < n; i++{
			ch <- buffer
		}
		time.Sleep(time.Second * 3)
	}

}
