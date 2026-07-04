package main

import "fmt"


func main (){
	var m[3][3] int
	var n int
	var soma_p int
	var soma_s int

	for i := 0; i< len(m);i++{
		for j :=0; j < len(m[0] );j++{
			n++
			m[i][j] = n

			
		}
	}
	for i:=0;i<len(m);i++{
		for j:=0; j<len(m[0]);j++{
			if i == j {
				soma_p += m[i][j]
			}
			if  j == len(m)-1-i{
				soma_s += m[i][j]
			}
			
		}
	}
	fmt.Println(soma_p)
	fmt.Println(soma_s)
	
	
}