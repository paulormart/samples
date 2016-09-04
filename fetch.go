package main

import (
	"fmt"
	"fetch"
)

func main(){

	a := fetch.Url("http://google.com")
	fmt.Println(a)
}
