package main

import "fmt"
import "encoding/json"

func main() {
	
	employee := Employee{123, "Bart", "Simpson", 2, true}
	
	b, err := json.Marshal(employee) // returns []byte containing the JSON data
	if ( err != nil ) {
		panic(err)
	}
	
	fmt.Printf("Marshal --> JSON : %s \n", b)
	
	employee2 := Employee{}
	err = json.Unmarshal(b, &employee2)
	if ( err != nil ) {
		panic(err)
	}
	fmt.Printf("JSON --> Unmarshal : %s \n", b)
	fmt.Printf("JSON --> Unmarshal : %+v \n", employee2)
	fmt.Printf("JSON --> Unmarshal : %v \n", employee2)
}