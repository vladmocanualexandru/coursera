 package main
 
 import (
	 "fmt"
	// "strconv"
 )
 
 func main() {
	type Grades int
	
	const (
		A Grades = iota
		B
		C
		D
		F
	)
	
	const PI = 3.14
	
	// var student1 Grades
	// student1 = A
	
	
	// student2 := B
	
	// fmt.Printf(strconv.Itoa(student1))
	
	var myPi = PI
	
	fmt.Printf("%f", myPi)
 }