package main

import "fmt"

func inverterSlice(slice *[]int){
	var s = *slice 
	for i,j := 0, len(s)-1; i<j;i,j = i-1 ,j-1{
		s[i],s[j]

	}
}

func main (){
	
}