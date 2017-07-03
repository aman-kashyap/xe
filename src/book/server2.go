package main

import (
	"io"
	"net/http"
)

//in hello we have two arguments one is http.ResponseWriter and its
//corresponding response stream which is an interface
//2nd is '*http.Request' and its corresponding request
//"io.WriteString" is a helper function to let you write a string in a given writable stream
func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello server!!!")
}

//In the main function called http.HandleFunc from package 'net/http' to register
//another function to be the handle function, which is hello function here.
//hello function accepts two arguments. the 1st is string type pattern,
//which is route you want to matchand its the route path here.
//The second is func (ResponseWriter, *Request)
//next "http.ListenAndServe(":8000",nil)" is called to listen on localhost with port 8000
/*func main(){
	http.HandleFunc("/",hello)
	http.ListenAndServe(":8000",nil)
}*/

//here we use 'mux' to register the handle function instead of directly registering
//from 'net/http' package BECAUSE 'net/http' has a default *ServeMux inside the package
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	http.ListenAndServe(":8000", mux)
}
