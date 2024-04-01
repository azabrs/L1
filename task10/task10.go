package main

import "fmt"

func grouping(temperatures []float64) map[int][]float64{
	dict := make(map[int][]float64)
	for _, val := range(temperatures){
		dict[int(val / 10 * 10)] =  append(dict[int(val / 10 * 10)], val)
	}
	return dict
}

func main(){
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	fmt.Printf("Breakdown by groups %v\n", grouping(temperatures))
}