package main
import "fmt"

func main() {
	var fpNumber float32
	fmt.Printf("Please enter a floating point number: ")
	
	fmt.Scan(&fpNumber)
	
	fmt.Printf("Here is the truncated version: %d \nGood bye!", int32(fpNumber))
}