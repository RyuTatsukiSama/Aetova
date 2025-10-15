package beyond100lines

import (
	"fmt"
)

func Concurrency_main() {
	var c chan int = make(chan int)
	for i := 0; i < 5; i++ {
		go cookingGopher(i, c)
	}

	for i := 0; i < 5; i++ {
		gopherID := <-c
		fmt.Println("gopher", gopherID, "finished the dish")
	}
}

func cookingGopher(id int, c chan int) {
	fmt.Println("gopher", id, "started cooking")
	c <- id
}
