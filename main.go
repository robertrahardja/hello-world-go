package main

import (
	"fmt"
	"mine/catfact"
	"mine/templ"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		catFact, err := catfact.FetchCatFact()
		if err != nil {
			http.Error(w, fmt.Sprintf("Error fetching cat fact: %s", err), http.StatusInternalServerError)
			return
		}

		// Render HelloWorld component
		err = templ.HelloWorld().Render(r.Context(), w)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error rendering HelloWorld: %s", err), http.StatusInternalServerError)
			return
		}

		// Render CatFactDisplay component
		err = templ.CatFactDisplay(catFact).Render(r.Context(), w)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error rendering CatFactDisplay: %s", err), http.StatusInternalServerError)
			return
		}
	})

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
