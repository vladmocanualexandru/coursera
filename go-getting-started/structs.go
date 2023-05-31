package main
import "fmt"

type Person struct {
	name string
	addr string
	phone string

}
	

func main(){


	p1 := new(Person)
	
	p2 := &Person{name:"joe", addr:"1 street", phone:"123"}
	
	p1.name = "John"
	
	fmt.Println(p1.name)
	
	fmt.Println(p2.name)
	fmt.Println(p2.addr)
	
}