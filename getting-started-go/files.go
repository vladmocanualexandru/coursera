package main

import (
	"fmt"
	"io/ioutil"
	"time"
)


func main(){

	dat, err := ioutil.ReadFile("test.txt")
	
	if err != nil {
		panic(err)
	}
	
	fmt.Println(string(dat))
	
	err = ioutil.WriteFile("test.txt", []byte(time.Now().String()), 0777)
	
}

