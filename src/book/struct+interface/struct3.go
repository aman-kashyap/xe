package main

import "fmt"

/*type Person struct{
	Name string
}
func(p Person) Talk(){
	fmt.Println("Hi, my name is", p.Name)
}
func main(){
	p:=Person{Name:"mishra"}
	p.Talk()
}*/
type Android struct{
	//Person
	Model string
}
func(a Android) Talk(){
	fmt.Println("hello!!! I don't have", a.Model)
}
func main(){
	a:=Android{Model:"android Nougat"}
	a.Talk()
}
