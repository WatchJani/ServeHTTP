package main

import (
	"fmt"
	"net/http"
)

type Router struct {
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Println("req")

	fmt.Println(req.URL.Path)

	// Route not found
	// http.NotFound(w, req)
}

func main() {
	router := &Router{}

	http.ListenAndServe(":5000", router)
}
