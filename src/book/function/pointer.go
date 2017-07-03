package main

import"fmt"

/*func zero(x *int){
	*x=0
}
func main(){
	x:=5
	zero(&x)
	fmt.Println(x)
}*/
/*func one(xPtr *int){
	*xPtr=66
}
func main(){
	xPtr :=new(int)
	one(xPtr)
	fmt.Println(*xPtr)
}*/
/*func square(x *float64) {
	*x = *x * *x
}
func main() {
	x := 1.51
	square(&x)
	fmt.Println(x)
}*/
func swap(x,y *int){
	*x=1
	*y=2
	*x,*y=*y,*x
}
func main(){
	x:=new(int)
	y:=new(int)
	//fmt.Println(x,y)
	swap(x,y)
	fmt.Println(*x,*y)
}
/*func swap(x,y *int){
	*x,*y=*y,*x
}
func main(){
	x:=1
	y:=2
	fmt.Println(x,y)
		swap(&x,&y)
	fmt.Println(x,y)
}*/