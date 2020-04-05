package api

import (
	// "fmt"
	"net/http"
	"strings"
)

// ECPHandler !
func ECPHandler(w http.ResponseWriter, r *http.Request) {
	// r.URL.Path creates a new path called /api
	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/ecpapi")
}
