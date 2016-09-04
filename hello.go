// https://www.youtube.com/watch?v=CF9S4QZuV30

package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	fmt.Println("Hello World")
	// terminal arguments
	if len(os.Args) != 2 {
		fmt.Println("go run hello.go hello world", os.Args[1])
		os.Exit(1)

	}
	fmt.Println("Job done!", os.Args[1])

	// for loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	// if statement
	age := 20
	if age >= 18 {
		fmt.Println("You can drive!")
	} else if age >= 18 {
		fmt.Println("You can vote!")
	} else {
		fmt.Println("Just relax!")
	}
	// switch statement
	switch age {
	case 16:
		fmt.Println("You can drive!")
	case 18:
		fmt.Println("You can vote!")
	default:
		fmt.Println("Have fun")
	}
	// arrays
	var a [3]float64
	a[0] = 1
	a[1] = 1.5
	a[2] = 112321.43
	fmt.Println(a[2])

	b := [3]float64{2, 4, 434.65}
	for i, value := range b {
		fmt.Println(value, i)
	}
	// slice
	s := []int{5, 4, 3, 2, 1}
	fmt.Println("Array s =", s)
	fmt.Println("Slice s[3:5] =", s[3:5])
	fmt.Println("Slice s[:3] =", s[:3])
	fmt.Println("Slice s[3:] =", s[3:])

	// start an array whith make function
	sm := make([]int, 5, 10)
	copy(sm, s)
	fmt.Println("Array via make", sm)
	sm = append(sm, 0, -1)
	fmt.Println("Array after append", sm)

	// maps
	presAge := make(map[string]int)
	presAge["Obama"] = 50
	presAge["Sanders"] = 70
	fmt.Println("Map", presAge, len(presAge))
	delete(presAge, "Sanders")
	fmt.Println("Map", presAge, len(presAge))

	// functions
	nums := []float64{1, 2, 3, 4, 5}
	fmt.Println("function add", add(nums))

	// closure
	num3 := 3
	double_n := func() int {
		num3 *= 2
		return num3
	}
	fmt.Println("closure", double_n())

	// recursion
	fmt.Println("recursion factorial", factorial(10))

	// defer can be used to call a function in the end of a function
	defer print2()
	print1()

	fmt.Println(safe_div(3, 0))
	fmt.Println(safe_div(3, 3))

	// pointers
	x := 0
	change_x(x)
	fmt.Println("change_x", x)
	change_x_now(&x)
	fmt.Println("change_x_now x=", x)
	fmt.Println("outside memory address x=", &x)

	// structs and interfaces
	rect := Rectangle{20, 50}
	fmt.Println("Rectangle =", rect.width, "wide")
	fmt.Println("Rectangle area =", rect.area())

	circ := Circle{4}
	fmt.Println("Rectangle area =", get_area(rect))
	fmt.Println("Circle area =", get_area(circ))

}

func add(numbers []float64) float64 {
	sum := 0.0
	for _, value := range numbers {
		sum += value
	}
	return sum
}

func factorial(num int) int {
	if num == 0 {
		return 1
	}
	return num * factorial(num-1)
}

func print2() {
	fmt.Println("print2")
}

func print1() {
	fmt.Println("print1")
}

func safe_div(num1, num2 int) int {
	defer func() {
		fmt.Println(recover())
	}()

	solution := num1 / num2
	return solution
}

func change_x(x int) {
	x = 4
	fmt.Println("inside change_x x=", x)
}

func change_x_now(x *int) {
	*x = 4
	fmt.Println("memory address inside change_x_now x=", x)
}

type Rectangle struct {
	height float64
	width  float64
}

func (r Rectangle) area() float64 {
	return r.height * r.width
}

type Shape interface {
	area() float64
}

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func get_area(s Shape) float64 {
	return s.area()
}
