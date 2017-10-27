package webserver

import (
	"net/http"
)

func Server(port int) {
	http.HandleFunc("/new", New)
	http.HandleFunc("/move", Move)
	http.HandleFunc("/suggest", Suggest)
	http.ListenAndServe(":2000", nil)
}
