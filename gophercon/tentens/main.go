package main
import(
"fmt"
"time"
)

func main(){
// STARTLOOP OMIT
	for i := 0; i <= 9; i++ {
		go func() {
			fmt.Println(i)
		}()
	}
// STOPLOOP OMIT
	time.Sleep(100*time.Millisecond)
}
