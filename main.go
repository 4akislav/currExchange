package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	fmt.Println("Server started on :8090")

	mux.HandleFunc("/usd/", func(w http.ResponseWriter, r *http.Request) {
		resp, err := http.Get("https://api.privatbank.ua/p24api/pubinfo?json&exchange&coursid=5")
		if err != nil {
			http.Error(w, "Failed to fetch data from API", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, "Failed to read API response body", http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "Response body: %s\n", body)

	})

	http.ListenAndServe("localhost:8090", mux)
}
