package main
import(
	"fmt"
	"sync"
	"net/http"
)


func Worker(wg1 *sync.WaitGroup,jobs <-chan string, result chan<- string){
	defer wg1.Done()
	for arg := range jobs {
		Http,_ := http.Get(arg)
		result <- arg
	}
}

func main() {
	jobs := make(chan string,10)
	result := make(chan string,10)
	var wg sync.WaitGroup
	wg.Add(3)
	go Worker(&wg, jobs,result)
	go Worker(&wg,jobs,result)
	go Worker(&wg,jobs,result)

	List []string
	for i := range List{
		jobs <- j
	}
	close(jobs)
	wg.Wait()
	close(result)

	for i := range result{
		fmt.Printf("%v\n",i)
	}


}
