package main

import (
	"context"
	"fmt"
	"net/http"
)

type Router struct {
	route map[string]http.HandlerFunc
}

func New() *Router {
	return &Router{
		route: make(map[string]http.HandlerFunc),
	}
}

func (r *Router) MyHandler(path string, handler http.HandlerFunc) {
	r.route[path] = handler
}

func Req(w http.ResponseWriter, r *http.Request) {
	// params := r.Context().Value("params")
	params := Params(r)

	fmt.Println(params)
}

func Params(r *http.Request) map[string]string {
	params := make(map[string]string)

	r.Context().Value("params")

	params["top"] = "top"

	return params
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if fn, ok := r.route[req.URL.Path]; ok {
		ctx := context.WithValue(req.Context(), "params", "janko")
		fn(w, req.WithContext(ctx))
	}

	fmt.Println("ne postoji")
}

func main() {
	router := New()

	router.MyHandler("/user", Req)

	http.ListenAndServe(":5000", router)
}
