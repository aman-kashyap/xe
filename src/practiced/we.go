package main

import (
	"fmt"
)

type Missile interface {
	Range() int
}
type Agni5 struct {
}

func (a Agni5) Range() int {
	return 5000
}

type Barak8 struct {
}

func (i Barak8) Range() int {
	return 90
}

type Nirbhay struct {
}

func (n Nirbhay) Range() int {
	return 3500
}

type Brahmos struct {
}

func (b Brahmos) Range() int {
	return 300
}

func main() {
	missiles := []Missile{Agni5{}, Barak8{}, Nirbhay{}, Brahmos{}}
	for _, missile := range missiles {
		fmt.Println(missile.Range())
	}
}
