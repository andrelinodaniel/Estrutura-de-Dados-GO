package main

import "fmt"

func main() {

	var sl []int
	var index int
	sl = []int{1, 2, 3}
	sl = append(sl[:index], sl[:index+1]...)
	fmt.Println(sl)
	

}
