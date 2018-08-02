package main

import(
	"fmt"
)

func main(){
// START OMIT
	freeVar := "Hello "
	f := func(s string){
		fmt.Println(freeVar + s)
	}
	f("Closures")
	freeVar = "Goodbye "
	f("Closures")
// STOP OMIT
}
