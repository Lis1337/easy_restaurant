package web

import (
	"fmt"
	categoryHandler "github.com/Lis1337/easy_restaurant/src/handler"
	"net/http"
)

func init() {
	// todo use https://pkg.go.dev/net/http#ServeMux
	getIndex()
	getCategory()

	http.ListenAndServe(":8000", nil)
}

func getIndex() {
	http.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello, Docker with net/http!")
		},
	)
}

func getCategory() {
	http.HandleFunc("/categories", categoryHandler.GetAllCategories)
	http.HandleFunc("/categories/{id}", categoryHandler.GetCategoryById)
}
