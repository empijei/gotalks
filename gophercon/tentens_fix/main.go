package main
import(
	"fmt"
	"time"
)

func main(){
// STARTLOOP OMIT
	for i := 0; i <= 9; i++ {
		i := i
		go func() {
			fmt.Println(i)
		}()
	}
// STOPLOOP OMIT
	time.Sleep(500*time.Millisecond)
}
