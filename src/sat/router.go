//this will serve static file with mux
package main

import (
	"flag"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	//	"time"
)

func main() {
	var dir string

	flag.StringVar(&dir, "dir", ".", "the directory to serve files from. Defaults to the current dir")
	flag.Parse()
	r := mux.NewRouter()

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(dir))))

	//r.host(www.example1.com)

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8020",

		//WriteTimeout: when net/http.ResponseWriter is trying to write something to the socket
		//and the write doesn't complete or doesn't manage to write at least one byte within the WriteTimeout
		//WriteTimeout: 15 * time.Second,

		//ReadTimeout: when net/http.Server is actively reading the next request from the socket
		//it times out if the read doesn't complete within the timeout
		//ReadTimeout:  15 * time.Second,

	}

	//ListenAndServe starts an HTTP server with a given address and handler.
	//The Handler is usually nil, which means to use Default Server mux Handle
	//Handle and HandleFunc add handlers to DefaultServeMux
	log.Fatal(srv.ListenAndServe())
}
