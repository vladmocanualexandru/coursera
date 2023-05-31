package main
import "fmt"

func main(){
	myArray := [...]int {10,20,30}
	
	for i,v := range myArray {
		fmt.Printf("%d:%d\n", i, v)
	}
}