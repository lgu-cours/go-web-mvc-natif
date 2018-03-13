package foo_test // _test by convention
// package foo // it works 

import (
	"testing"
	
	"fmt"
	"strconv"
)

func Test1(t * testing.T ) {
	    fmt.Println(t.Name()) // prints function name
	    
	    int1 := 123 
	    s := strconv.Itoa(int1)
	    int2, err := strconv.Atoi(s)
	    if ( err != nil ) {
	    	t.Error("Atoi error")
	    }
	    
	    if ( int1 != int2 ) {
	    	t.Error("int1 != int2")
	    }
	    t.Name()
}

func Test2(t * testing.T ) {
	    fmt.Println(t.Name()) // prints function name
	    t.Fail()  // test failed without message 
}

func Test3(t * testing.T ) {
	    fmt.Println(t.Name()) // prints function name
	    t.Error("My test error") // test failed with error message 
}

func Foo(t * testing.T ) { // ignored (not starting with "Test")
	    fmt.Println(t.Name()) // prints function name
}

func TestFoo(t * testing.T ) { // processed
	    fmt.Println(t.Name()) // prints function name
}
