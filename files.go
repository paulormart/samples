// https://www.youtube.com/watch?v=CF9S4QZuV30

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	// "strconv"
)

func main() {
	file, err := os.Create("sample.txt")
	if err != nil {
		log.Fatal(err)
	}

	file.WriteString("Hello World")
	file.Close()
	stream, err := ioutil.ReadFile("sample.txt")
	if err != nil {
		log.Fatal(err)
	}

	read_string := string(stream)
	fmt.Println("ReadFile =", read_string)

}
