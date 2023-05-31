package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("This program generates a JSON based on the information provided by the user.")

	// I had to use a bufio.NewReader because normal Scan methods would only read the first word
	reader := bufio.NewReader(os.Stdin)

	// Read the name
	fmt.Printf("Name: ")
	name, _ := reader.ReadString('\n')
	name = strings.Replace(name, "\r\n", "", -1)

	// Read the address
	fmt.Printf("Address: ")
	address, _ := reader.ReadString('\n')
	address = strings.Replace(address, "\r\n", "", -1)

	// Create the map
	userMap := map[string]string{
		"name":    name,
		"address": address,
	}

	// Convert the map to bytes array representing the json string
	userMapBytes, err := json.Marshal(userMap)

	if err != nil {
		fmt.Println("Unable to marshall the user map.")
		panic(err)
	}

	// Create a string using the bytes array and display it in the console
	userMapString := string(userMapBytes)
	fmt.Printf("%s\nGood bye!", userMapString)

}
