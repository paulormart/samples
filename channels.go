package main

import (
	"fmt"
	"strconv"
	"time"
)

var pizza_num = 0
var pizza_name = ""

func makeDough(channel chan string) {
	pizza_num++
	pizza_name = "Pizza #" + strconv.Itoa(pizza_num)
	fmt.Println("Make dough and send for sauce for", pizza_name)
	// pass to the sauce channel
	channel <- pizza_name

	time.Sleep(time.Millisecond * 100)
}

func addSauce(channel chan string) {

	pizza := <-channel
	fmt.Println("Add sauce and send", pizza, "for toppings")
	channel <- pizza

	time.Sleep(time.Millisecond * 100)
}

func addToppings(channel chan string) {
	pizza := <-channel
	fmt.Println("Add toppings to", pizza, "and ship")
	time.Sleep(time.Millisecond * 100)
}

func main() {

	string_chan := make(chan string)

	for i := 0; i < 1; i++ {
		fmt.Println("Start", string_chan)
		go makeDough(string_chan)
		go addSauce(string_chan)
		go addToppings(string_chan)
	}
	time.Sleep(time.Millisecond * 1000000)
}
