package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x, y, r float64
}
type Rectangle struct {
	x1, y1, x2, y2 float64
}
type Shape interface {
	area() float64
	//perimeter() float64
}

func (c Circle) area() float64 {
	return math.Pi * c.r * c.r
}
func (r Rectangle) distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}
func (r Rectangle) area() float64 {
	l := r.distance(r.x1, r.y1, r.x1, r.y2)
	w := r.distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}
func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}
func main() {
	c := Circle{0, 0, 5}
	r := Rectangle{0, 0, 10, 10}
	shapes := []Shape{r, c}
	for _, shape := range shapes {
		fmt.Println(shape.area())
	}
	//fmt.Println(totalArea(&c, &r))
}

/*
func (c Circle) perimeter() float64{
	return math.Pi*2*c.r
}
func (r Rectangle) distance(x1,y1,x2,y2 float64) float64{
	a:=x2-x1
	b:=y2-y1
	return math.Sqrt(a*a + b*b)
}
func (r Rectangle) perimeter() float64{
	l:=r.distance(r.x1,r.y1,r.x1,r.y2)
	w:=r.distance(r.x1,r.y1,r.x2,r.y1)
	return 2*(l+w)
	}
func totalPerimeter(shapes ...Shape) float64{
	var perimeter float64
	for _, s:=range shapes{
		perimeter+=s.perimeter()
	}
	return perimeter
}
func main(){
	c:=Circle{0,0,5}
	r:=Rectangle{0,0,10,10}
	shapes:=[]Shape{r,c}
	for _,shape:=range shapes{
		fmt.Println(shape.perimeter())
	}
	fmt.Println(c.perimeter(),r.perimeter())
	fmt.Println(totalPerimeter(&c,&r))
}*/
