// http://openmymind.net/assets/go/go.pdf
// Channels page 76
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	c := make(chan int)
	for i := 0; i < 5; i++ {
		worker := &Worker{id: i}
		go worker.process(c)
	}

	for {
		select {
			case c <- rand.Int():
		  	//optional code here
			case t := <-time.After(time.Millisecond * 100):
		    fmt.Println("timed out at ", t)
		  // default:
		  //   //this can be left empty to silently drop the data
		  //   fmt.Println("dropped")
		}
		// fmt.Println(len(c))
		time.Sleep(time.Millisecond * 50)
	}
}

type Worker struct {
	id int
}

func (w *Worker) process(c chan int) {
	for {
		data := <-c
		fmt.Printf("worker %d got %d\n", w.id, data)
		time.Sleep(time.Second)
		// select {
		// case data := <-c:
		// 	fmt.Printf("worker %d got %d\n", w.id, data)
		// case <-time.After(time.Millisecond * 10):
		// 	fmt.Printf("worker %d has some break time\n", w.id)
		// 	time.Sleep(time.Second)
		// }
	}
}
