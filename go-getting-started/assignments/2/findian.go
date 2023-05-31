package main
import (
	"fmt"
	"strings"
)

const PREFIX = "i"
const SUBSTR = "a"
const SUFFIX = "n"

func main() {
	var userString string
	
	fmt.Printf("Please enter a string: ")
	fmt.Scan(&userString)
	
	// in order to ignore case, transform user string to lowercase version
	userString = strings.ToLower(userString)
	
	// 3 part conditions check
	foundIan := strings.HasPrefix(userString, PREFIX)
	foundIan = foundIan && strings.Contains(userString, SUBSTR)
	foundIan = foundIan && strings.HasSuffix(userString, SUFFIX)
	
	// output according to check above
	if foundIan {
		fmt.Printf("Found!")
	} else {
		fmt.Printf("Not found!")
	}
}