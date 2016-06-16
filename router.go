package gaufre

import (
	"strings"
	"net/http"
	"log"
	"time"
	"strconv"
)

type Router struct {
	Routes map[string]Route
}

func (r *Router) NewRoute(method string, path string, action func(http.ResponseWriter, *http.Request)) {
	method = strings.ToUpper(method)
	r.Routes[strings.Join([]string{method, path}, "")] = Route{method, path, action}
}

func (r *Router) Get(path string, action func(http.ResponseWriter, *http.Request)) {
	r.NewRoute("GET", path, action)
}

func (r *Router) Post(path string, action func(http.ResponseWriter, *http.Request)) {
	r.NewRoute("POST", path, action)
}

func (r *Router) Put(path string, action func(http.ResponseWriter, *http.Request)) {
	r.NewRoute("PUT", path, action)
}

func (r *Router) Patch(path string, action func(http.ResponseWriter, *http.Request)) {
	r.NewRoute("PATCH", path, action)
}

func (r *Router) Delete(path string, action func(http.ResponseWriter, *http.Request)) {
	r.NewRoute("DELETE", path, action)
}

func (r *Router) Listen(port int) {
	var server = &http.Server{
		Addr:           strings.Join([]string{":", strconv.Itoa(port)}, ""),
		Handler:        &Handler{},
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(server.ListenAndServe())
}

func GetRouter() *Router {
	return router
}
