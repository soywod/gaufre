package gaufre

import "net/http"

type Route struct {
	method string
	path   string
	action func(http.ResponseWriter, *http.Request)
}
