package main

import "fmt"

type Tank interface {
	Area() float32
	Volume() float32
}
type Value struct {
	Radius float32
	Height float32
}

func (v Value) Area() float32 {
	return (2 * 3.41 * v.Radius * v.Radius) + (2 * v.Radius * v.Height)
}
func (v Value) Volume() float32 {
	return (3.14 * v.Radius * v.Radius * v.Height)
}
func main() {
	val1 := Value{60.5, 130.5}
	val2 := Value{62.5, 128.5}
	fmt.Println("area of val1 is :", val1.Area())
	fmt.Println("area of val2 is :", val2.Area())
	fmt.Println("volume of val1 is :", val1.Volume())
	fmt.Println("volume of val1 is :", val2.Volume())
	fmt.Println(":::", Tank.Area(val1))
	

}
