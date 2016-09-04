// https://www.youtube.com/watch?v=CF9S4QZuV30

package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	hello := "Hello World"
	csv_hello := "H,e,l,l,o,W,o,r,l,d"
	fmt.Println(strings.Contains(hello, "He"))
	fmt.Println(strings.Index(hello, "Wo"))
	fmt.Println(strings.Count(hello, "o"))
	fmt.Println(strings.Replace(hello, "o", "a", 2))
	fmt.Println(strings.Split(csv_hello, ","))
	// sort
	l := strings.Split(csv_hello, ",")
	sort.Strings(l)
	fmt.Println(l)
	//
	fmt.Println(strings.Join(l, "+ "))
	// conversions
	rand_int := 5
	rand_float := 5.5
	rand_string := "200"
	rand_string2 := "200.6"
	fmt.Println(float64(rand_int))
	fmt.Println(int(rand_float))
	new_int, _ := strconv.ParseInt(rand_string, 0, 64)
	fmt.Println(new_int)
	new_float, _ := strconv.ParseFloat(rand_string2, 64)
	fmt.Println(new_float)
}
