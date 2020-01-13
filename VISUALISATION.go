package main

import (
	"fmt"
	"net/http"
)

// ////////////////////////////////////
type Visualizator struct {
	ID     int
	Name   string
	secret string
}
func (m Visualizator) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	fmt.Println(req.Body, " \n###\n ", req.Header, " \n###\n ", req.Method, " \n###\n ", req.Host, " \n###\n ", req.RemoteAddr)
	fmt.Fprint(resp, "Hi man: "+string(m.ID)+"!")
}
// ////////////////////////////////////

// /////// SERVER FUNCTIONS //////////
func ServeHello(resp http.ResponseWriter, req *http.Request) {
	fmt.Println(req.Body, " \n###\n ", req.Header, " \n###\n ", req.Method, " \n###\n ", req.Host, " \n###\n ", req.RemoteAddr)
	fmt.Fprint(resp, "Hi man!")
}
func ServeDescr(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprint(resp, "I am very strange function for fun!")
}
// ////////////////////////////////////



func main() {
	//visHandler := Visualizator{ID: 1,
	//	Name:   "Monster",
	//	secret: "Ohohohho"}
	fmt.Println("Server is listening...")

	http.HandleFunc("/hello", ServeHello)
	http.HandleFunc("/descr", ServeDescr)

	http.ListenAndServe("192.168.224.108:50509", nil)
}
