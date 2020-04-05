package api

import (
	"net/http"
	"strings"
)

// APIHandler !
func APIHandler(w http.ResponseWriter, r *http.Request) {
	// r.URL.Path creates a new path called /api

	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/api")
	if strings.HasPrefix(r.URL.Path, "/tokens") {
		ECPHandler(w, r)
		return
	}

}
