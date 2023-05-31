package main

import (
	"fmt"
	"encoding/json"
)

type Person struct {
	Name string
	Age int
}

func main(){

	p1 := Person{Name:"Joe", Age:69}
	p1Bytes, err := json.Marshal(p1)
	
	if err != nil {
		panic(err)
	}
	
	p1Str := string(p1Bytes)
	fmt.Println(p1Str)
	
	p2Str := "{\"Name\":\"Little Jimmy\", \"Age\":10}"
	
	var p2 Person
	err = json.Unmarshal([]byte(p2Str), &p2)
	
	fmt.Printf("%s %dyo\n", p2.Name, p2.Age)
	
}

