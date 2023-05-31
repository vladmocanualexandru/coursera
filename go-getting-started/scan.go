package main

import (
	"fmt"
	"math/rand"
)

func main(){

	goal := rand.Intn(100)
	
	fmt.Printf("Let's go! Guess my number!\n")
	
	gameover := false
	for gameover == false {
		var userChoice int
		fmt.Printf("Guess:")
		fmt.Scan(&userChoice)
		
		switch {
			case userChoice < goal:
				fmt.Printf("Too low... \n")
			case userChoice > goal:
				fmt.Printf("Too high... \n")
			case userChoice == goal: {
				fmt.Printf("Correct! \n")
				gameover = true
			}
				
		}
	}
	
}