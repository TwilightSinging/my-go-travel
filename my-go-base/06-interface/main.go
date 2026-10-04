package main

import (
	"fmt"
	"math"
)

// GO没有函数重载但是有函数重写和多态
type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	Radius float64
}

// circle area
func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

// circle perimeter
func (c Circle) Perimeter() float64 {
	return 3.14 * 2 * c.Radius
}

// rectangle
type Rectangle struct {
	Width  float64
	Height float64
}

// Area of rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// perimeter of rectangle
func (r Rectangle) Perimeter() float64 {
	return r.Width * r.Height
}

// triangle
type Triangle struct {
	A float64
	B float64
	C float64
}

// area of triangle
func (t Triangle) Area() float64 {
	//海伦公式
	//先求周长
	s := (t.A + t.B + t.C) / 2
	return math.Sqrt(s * (s - t.A) * (s - t.B) * (s - t.C))
}

// perimeter of triangle
func (t Triangle) Perimeter() float64 {
	return t.A + t.B + t.C
}
func printShapeInfo(s Shape) {
	switch v := s.(type) {
	case Circle:
		fmt.Printf("圆形: 半径=%.2f, 面积=%.2f,周长=%.2f\n", v.Radius, v.Area(), v.Perimeter())
	case Rectangle:
		fmt.Printf("矩形: 宽=%.2f, 高=%.2f, 面积=%.2f,周长=%.2f\n", v.Width, v.Height, v.Area(), v.Perimeter())
	case Triangle:
		fmt.Printf("三角形: 边长=%.2f,%.2f,%.2f, 面积=%.2f,周长=%.2f\n", v.A, v.B, v.C, v.Area(), v.Perimeter())
	default:
		fmt.Printf("未知形状: %T\n", v)
	}
}

func main() {
	var s Shape
	var c Circle = Circle{
		Radius: 1,
	}
	s = c
	printShapeInfo(s)
	var r Rectangle = Rectangle{
		Width:  1,
		Height: 1,
	}
	s = r
	printShapeInfo(s)
	var t Triangle = Triangle{
		A: 3,
		B: 4,
		C: 5,
	}
	s = t
	printShapeInfo(s)
}
