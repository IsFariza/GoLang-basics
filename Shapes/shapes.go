package Shapes

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

type Triangle struct {
	Side1, Side2, Side3 float64
}

func (t Triangle) Area() float64 {
	s := t.Perimeter() / 2 //semi perimeter
	return math.Sqrt(s * (s - t.Side1) * (s - t.Side2) * (s - t.Side3))
}
func (t Triangle) Perimeter() float64 {
	return t.Side1 + t.Side2 + t.Side3
}

type Square struct {
	side float64
}

func (s Square) Area() float64 {
	return s.side * s.side
}
func (s Square) Perimeter() float64 {
	return s.side * 4
}

func CreateShapes() []Shape {
	shapes := []Shape{
		&Rectangle{5.5, 12},
		&Circle{16},
		&Square{23},
		&Triangle{8, 9.7, 5},
	}
	for _, shape := range shapes {
		switch shape := shape.(type) {
		case *Rectangle:
			fmt.Printf("- Rectangle: width = %v, height = %v\n", shape.Width, shape.Height)
		case *Circle:
			fmt.Printf("- Circle: radius = %v\n", shape.Radius)
		case *Square:
			fmt.Printf("- Square: side = %v\n", shape.side)
		case *Triangle:
			fmt.Printf("- Triangle: side1 = %v, side2 = %v, side3 = %v\n", shape.Side1, shape.Side2, shape.Side3)
		}
	}
	return shapes
}

func CreateCustomShapes() []Shape {
	shapes := make([]Shape, 0)

	fmt.Println("Enter parameters for Rectangle: ")
	fmt.Println("Width: ")
	var width float64
	fmt.Scan(&width)
	fmt.Println("Height: ")
	var height float64
	fmt.Scan(&height)
	shapes = append(shapes, &Rectangle{width, height})

	fmt.Println("Enter parameters for Circle: ")
	fmt.Println("Radius: ")
	var radius float64
	fmt.Scan(&radius)
	shapes = append(shapes, &Circle{radius})

	var side1 float64
	var side2 float64
	var side3 float64
	fmt.Println("Enter parameters for Triangle: ")
	fmt.Println("Side 1: ")
	fmt.Scan(&side1)
	fmt.Println("Side 2: ")
	fmt.Scan(&side2)
	fmt.Println("Side 3: ")
	fmt.Scan(&side3)
	shapes = append(shapes, &Triangle{side1, side2, side3})

	var side float64
	fmt.Println("Enter parameters for Square: ")
	fmt.Println("Side: ")
	fmt.Scan(&side)
	shapes = append(shapes, &Square{side})

	return shapes
}
func IterateShapes(shapes []Shape) {
	fmt.Println("\nArea: ")
	for _, shape := range shapes {
		switch shape := shape.(type) {
		case *Rectangle:
			fmt.Printf("- Rectangle: area: %.2f, perimeter: %.2f\n", shape.Area(), shape.Perimeter())
		case *Circle:

			fmt.Printf("- Circle: area: %.2f, perimeter: %.2f\n", shape.Area(), shape.Perimeter())
		case *Square:

			fmt.Printf("- Square: area: %.2f, perimeter: %.2f\n", shape.Area(), shape.Perimeter())
		case *Triangle:
			fmt.Printf("- Triangle: area: %.2f, perimeter: %.2f\n", shape.Area(), shape.Perimeter())
		default:
			fmt.Printf("- Unknown shape")
		}
	}
}
