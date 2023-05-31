package main

import "fmt"

func main(){

	var a int = 10
	var b *int = &a
	
	a = a+10
	
	(*b) = a + 20
	
	fmt.Printf("%d", *b)

}
