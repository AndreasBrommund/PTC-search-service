package controller

import (
	"net/http"
)

func Client(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/client", http.StatusMovedPermanently) // http.StatusMovedPermanently = 301
}
