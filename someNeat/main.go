package main

import (
	"fmt"
)

const (
	PI = 3.1415926
)

// Define the Shape interface
type Shape interface {
	Area() float64
}

// Rectangle struct
type Rectangle struct {
	Width  float64
	Height float64
}

// Circle struct
type Circle struct {
	Radius float64
}

// Implement the Area method for Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Implement the Area method for Circle
func (c Circle) Area() float64 {
	return PI * c.Radius * c.Radius
}

// Function to print the area of any Shape
func PrintArea(s Shape) {
	fmt.Printf("Area: %.2f\n", s.Area())
}

func getEmployee() (string, int, float32) {
	return "Bill", 50, 6789.50
}

func main() {

	rectangle := Rectangle{Width: 5, Height: 3}
	circle := Circle{Radius: 2.5}

	// Call PrintArea on rectangle and circle, both implement the Shape interface
	PrintArea(rectangle) // Prints the area of the rectangle
	PrintArea(circle)    // Prints the area of the circle

	// MySlice is a Slice of a Shape array with the length of 1
	MySlice := make([]Shape, 1)
	MySlice[0] = circle
	fmt.Println(MySlice, "Length:", len(MySlice))
	MySlice = append(MySlice, rectangle)
	fmt.Println(MySlice, "Length:", len(MySlice))

	name, age, salary := getEmployee()
	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(salary)

	//var number int = 1
	//// is the same as:
	//number1 := 1
	//
	//var message string = "Hello from Go"
	//// is the same as:
	//message1 := "Hello from Go"
	//fmt.Printf("Types of variables: %T ,%T ,%T ,%T", message, message1, number, number1)
}
