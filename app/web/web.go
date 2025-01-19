package web

import (
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, Docker with net/http!")
	})

	http.ListenAndServe(":8000", nil)
}
