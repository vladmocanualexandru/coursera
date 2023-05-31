package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type User struct {
	fname string
	lname string
}

func main() {

	// a brief description of the program + example of input file
	fmt.Println("This program will read from a file multiple names (one pair of first name/last name (separated by 1 space) per line). For each line, it will create a struct and add it to a slice. At the end, it will iterate through all the structs in the slice and display them on screen.")
	fmt.Printf("\nExample of file:\n\n")
	fmt.Println("Hafsa Ashley")
	fmt.Println("Brendon Watkins")
	fmt.Println("Rodney Adkins")
	fmt.Println("Sulayman Ramsey")
	fmt.Println("Omari Reilly")
	fmt.Println("Angus Carson")
	fmt.Println("Mohamad Burgess")

	// read file path from user
	var filePath string
	fmt.Printf("\nPath to file: ")
	fmt.Scan(&filePath)

	// open the connection to the file
	userDataFile, err := os.Open(filePath)

	// if something is wrong like the file not existing, display an appropriate message
	if err != nil {
		fmt.Println("Unable to open the provided file")
		panic(err)
	}

	// init the user slice
	userSlice := make([]User, 0)

	// init the scanner that will be used to read from the file
	fileScanner := bufio.NewScanner(userDataFile)

	// configure the separator
	fileScanner.Split(bufio.ScanLines)

	// while there are still things in the file to be read, the loop will be active
	for fileScanner.Scan() {
		// each line is split by space
		userDataLineTokens := strings.Split(fileScanner.Text(), " ")
		firstName := userDataLineTokens[0]
		lastName := userDataLineTokens[1]

		//a new User struct is appended to the slice
		userSlice = append(userSlice, User{fname: firstName, lname: lastName})
	}

	// once reading is finished, close the file connection
	userDataFile.Close()

	// display slice contents
	fmt.Println("Slice contains:")
	for key, user := range userSlice {
		fmt.Printf("%d %s %s\n", key, user.fname, user.lname)
	}

	fmt.Println("Have a nice day!")
}
