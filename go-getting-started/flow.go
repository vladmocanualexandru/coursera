package main

import "fmt"

func main(){

	// for loop
	for i:=0; i<10; i++ {
		
		// if statement
		if i%2==0 {
			j:=0
			
			// "while" loop
			for j<i {
				fmt.Printf("%d",  j*10)
				
				// switch with tag
				switch j {
					case 1:
						fmt.Printf("* ")
					case 3:
						fmt.Printf("@ ")
					case 5:
						fmt.Printf("# ")
					default:
						fmt.Printf(" ")
				}
				
				j++
			}
			fmt.Printf("\n")
			
			// infinite loop
			j=0
			for {
				// tagless switch
				switch {
					case j == 0 || j == 29:
						fmt.Printf("|")
					default:
						fmt.Printf("-")
				}
				
				j++
				
				if j==30 {
					break
				}
			}
			fmt.Printf("\n")
		}
	}
	
	
}