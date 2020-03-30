//Hands-on exercise #5
//create a type SQUARE
//create a type CIRCLE
//attach a method to each that calculates AREA and returns it
//circle area= π r^2
//square area = L * W
//create a type SHAPE that defines an interface as anything that has the AREA method
//create a func INFO which takes type shape and then prints the area
//create a value of type square
//create a value of type circle
//use func info to print the area of square
//use func info to print the area of circle

package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float32
}

type square struct {
	length float32
	width  float32
}

func (c circle) area() float32 {
	return c.radius * math.Pi * c.radius

}
func (s square) area() float32 {
	return s.length * s.width
}

type shape interface {
	area() float32
}

func info(s shape) {
	fmt.Println(s.area())
}
func main() {
	c := circle{radius: 14.14,}
	s := square{
		length: 10,
		width:  5,
	}

	info(c)
	info(s)
}
