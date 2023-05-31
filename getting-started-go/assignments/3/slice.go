package main

import (
	"fmt"
	"strings"
	"strconv"
)

const INITIAL_SLICE_SIZE = 3

func main(){
	// create the initial slice
	slice := make([]int, INITIAL_SLICE_SIZE)
	
	fmt.Println("Type an integer to add to the slice. Type X to quit.")
	
	userInput := ""
	
	// while user inputs anything besides X, loop is active
	for userInput != "x" {
		fmt.Printf("> ")
		fmt.Scan(&userInput)
		
		// lowercase input to simplify things
		userInput = strings.ToLower(userInput)
		
		if userInput != "x" {
			// input is not x, loop is not over
			
			// try to convert to int, take into account any potential errors
			intNumber, err := strconv.Atoi(userInput)
			
			if err != nil {
				// error is not nil -> input is not an integer
				fmt.Println("Doesn't look like an integer... Try again!")
			} else {
				// no error, input is integer; add to slice
				slice = append(slice, intNumber)
				
				// sort the slice using BubbleSort like in we did in highschool :D
				sorted:=false
				for !sorted {
					sorted = true
					
					for i:=0;i<len(slice)-1;i++ {
						if slice[i]>slice[i+1] {
							sorted = false
							
							temp := slice[i]
							slice[i] = slice[i+1]
							slice[i+1] = temp
						}
					}
				}
				
				// confirm integer added to slice and display slice contents
				fmt.Printf("Added to slice! Slice=(")
				
				for _,v := range slice {
					 fmt.Printf(" %d", v)
				}
				fmt.Printf(" )\n")
			}
		} else {
			// input is x, loop is over
			fmt.Println("Good bye!")
		}
		
	}
}