package main
import(
	"fmt"
	"runtime"
)

func main(){
	var i byte
// START OMIT
	go func() {
		for i = 0; i <= 255; i++ {
		}
	}()
	fmt.Println("Dropping mic")
	// Yield execution to force executing other goroutines
	runtime.Gosched()
	runtime.GC()
	fmt.Println("Done")
// STOP OMIT
}
