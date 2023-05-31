package main
import "fmt"

func main(){
	
	var idMap map[string]int
	idMap = make(map[string]int)
	idMap["jane"] = 123
	
	var idMap2 map[string]int = make(map[string]int)
	idMap2["jane"] = 456
	
	idMap3 := make(map[string]int)
	idMap3["jane"] = 789
	
	idMap4 := map[string]int {
		"jane":10101,
		"joe":1101,
		"robin":1111001,
	}
	
	fmt.Println(idMap4["jane"])
	
	_, p := idMap["joe"]
	
	if p {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
	
	_, p = idMap["jane"]
	
	if p {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
	
	for k,v := range idMap4 {
		fmt.Printf("%s - %d\n", k,v)
	}
	
}