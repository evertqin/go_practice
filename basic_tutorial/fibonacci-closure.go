package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
    prev := 0
    curr := 1
    idx := 0


    f := func() int {
	idx += 1
	if idx == 0 {
	    return 0
	} else if idx == 1 {
	    return 1
	} else {
	    next := prev + curr
	    prev = curr
	    curr = next
	    return curr
	}
    }

    return f
}

func main() {
    f := fibonacci()
    for i := 0; i < 10; i++ {
	fmt.Println(f())
    }
}