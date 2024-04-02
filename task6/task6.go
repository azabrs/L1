package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func  closeWithChan(ch chan int, wg *sync.WaitGroup){
	fmt.Println("goroutine started")
	<- ch
	fmt.Println("goroutine has finished")
	wg.Done()
}

func  closeWithCtx(ctx context.Context){
	fmt.Println("goroutine started")
	<- ctx.Done()
	fmt.Println("goroutine has finished")
}

func main(){
	// first solution using chan
	//ch := make(chan int)
	//var wg sync.WaitGroup
	//wg.Add(1)
	//go closeWithChan(ch, &wg)
	//time.Sleep(time.Second * 3)
	//ch <- 1
	//wg.Wait()


	// second solution using context
	ctx, cancel := context.WithCancel(context.Background())
	go closeWithCtx(ctx)

	time.Sleep(time.Second)
	cancel()
	time.Sleep(time.Second)
}