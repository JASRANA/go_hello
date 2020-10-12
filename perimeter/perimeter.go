package perimeter

import "math"

type Rectangle struct {
	Width float64
	Length float64
}

// 通过这个方式指明调用者
func (r Rectangle) Area() float64 {
	return r.Width * r.Length
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Triangle struct {
	Base float64
	Height float64
}

func (t Triangle) Area() float64 {
	return t.Base * t.Height / 2
}

// 接口隐式匹配
type Shape interface {
	Area() float64
}

func Perimeter(rectangle Rectangle) float64 {
	return (rectangle.Length + rectangle.Width) * 2
}

func Area(rectangle Rectangle) float64 {
	return rectangle.Length * rectangle.Width
}

