package gaufre

import (
	"strings"
	"net/http"
)

type Handler struct{}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var key = strings.Join([]string{r.Method, r.URL.Path}, "")

	if route, ok := router.Routes[key]; ok {
		route.action(w, r)
	}
}
