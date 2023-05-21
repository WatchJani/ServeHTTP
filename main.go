package main

import (
	"context"
	"fmt"
	"net/http"
	"strings"
)

type Router struct {
	routeEvent string
	route      map[string]http.HandlerFunc
}

func New() *Router {
	return &Router{
		route: make(map[string]http.HandlerFunc),
	}
}

func GetURL(path string) string {
	if strings.Contains(path, ":") {
		return strings.Split(path, ":")[0]
	}

	return path
}

func (r *Router) MyHandler(path string, handler http.HandlerFunc) {
	r.route[GetURL(path)], r.routeEvent = handler, path
}

func Req(w http.ResponseWriter, r *http.Request) {
	// params := r.Context().Value("params")
	params := Params(r)

	fmt.Println(params)
}

func Params(r *http.Request) map[string]string {
	params := make(map[string]string)

	r.Context().Value("params")

	return params
}

func DefaultURL(dynamicURL, reqURL string) string {
	dynamicURL = GetURL(dynamicURL)

	if strings.HasPrefix(reqURL, dynamicURL) {
		return dynamicURL
	}

	return ""
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	ourHandler := DefaultURL(r.routeEvent, req.URL.Path)

	if fn, ok := r.route[ourHandler]; ok {
		ctx := context.WithValue(req.Context(), "params", r.routeEvent)
		fn(w, req.WithContext(ctx))
		return
	}

	http.NotFound(w, req)
}

func main() {
	router := New()

	router.MyHandler("/user/:id", Req)

	http.ListenAndServe(":5000", router)
}
